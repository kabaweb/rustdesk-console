package api

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"rustdesk-nexus-server/config"
	gh "rustdesk-nexus-server/github"
	"rustdesk-nexus-server/store"

	"github.com/google/uuid"
)

type BuildHandler struct {
	cfg     *config.Config
	store   *store.Store
	gh      *gh.Client
	storage string
}

func NewBuildHandler(cfg *config.Config, s *store.Store, ghClient *gh.Client, storagePath string) *BuildHandler {
	return &BuildHandler{
		cfg:     cfg,
		store:   s,
		gh:      ghClient,
		storage: storagePath,
	}
}

type NexusGenerateRequest struct {
	OS        string                 `json:"os"`
	Arch      string                 `json:"arch"`
	Custom    map[string]interface{} `json:"custom"`
	InstallID string                 `json:"install_id"`
}

func (h *BuildHandler) Generate(w http.ResponseWriter, r *http.Request) {
	userGUID := r.Header.Get("X-User-GUID")

	existing, _ := h.store.FindActiveBuild(userGUID)
	if existing != nil {
		writeJSON(w, http.StatusConflict, map[string]string{"error": "already have an active build"})
		return
	}

	if h.cfg.MaxBuildsPerMonth > 0 {
		count, _ := h.store.CountBuildsThisMonth(userGUID)
		if count >= h.cfg.MaxBuildsPerMonth {
			writeJSON(w, http.StatusTooManyRequests, map[string]string{"error": "monthly build limit reached"})
			return
		}
	}

	var req NexusGenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.OS == "" {
		req.OS = "windows"
	}
	if req.Arch == "" {
		req.Arch = "x86_64"
	}

	appName := ""
	if req.Custom != nil {
		if n, ok := req.Custom["app-name"].(string); ok {
			appName = n
		}
	}

	customJSON, _ := json.Marshal(req.Custom)

	buildUUID := uuid.New().String()

	build := &store.Build{
		UUID:     buildUUID,
		UserGUID: userGUID,
		OS:       req.OS,
		Arch:     req.Arch,
		AppName:  appName,
		Custom:   string(customJSON),
		Status:   "pending",
		Files:    "[]",
		Message:  "",
	}

	if err := h.store.CreateBuild(build); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create build record"})
		return
	}

	if h.gh != nil && h.cfg.GitHubToken != "" {
		go h.triggerGitHubBuild(build, req)
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"uuid":    buildUUID,
		"status":  "pending",
		"message": "Build submitted",
	})
}

func (h *BuildHandler) triggerGitHubBuild(build *store.Build, req NexusGenerateRequest) {
	customJSON, _ := json.Marshal(req.Custom)

	input := &gh.WorkflowDispatchInput{
		RequestID:    build.UUID,
		RustDeskRepo: h.cfg.RustDeskRepo,
		RustDeskRef:  h.cfg.RustDeskRef,
		AppName:      build.AppName,
		CustomConfig: string(customJSON),
		OS:           req.OS,
		Arch:         req.Arch,
	}

	if err := h.gh.DispatchWorkflow(input); err != nil {
		h.store.UpdateBuild(build.UUID, "failed", "[]", fmt.Sprintf("dispatch: %v", err), 0)
		return
	}

	time.Sleep(5 * time.Second)

	run, err := h.gh.GetLatestRunForRequest(build.UUID)
	if err != nil || run == nil {
		h.store.UpdateBuild(build.UUID, "building", "[]", "", 0)
		return
	}

	h.store.UpdateBuild(build.UUID, "building", "[]", "", run.ID)
}

func (h *BuildHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	uuid := extractUUID(r.URL.Path, "/v1/client/generate/")

	build, err := h.store.GetBuild(uuid)
	if err != nil || build == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "build not found"})
		return
	}

	files := []string{}
	json.Unmarshal([]byte(build.Files), &files)
	if files == nil {
		files = []string{}
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"uuid":    build.UUID,
		"status":  build.Status,
		"files":   files,
		"message": build.Message,
	})
}

func (h *BuildHandler) Download(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v1/client/download/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid path"})
		return
	}
	uuid := parts[0]
	filename := parts[1]

	filePath := filepath.Join(h.storage, uuid, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "file not found"})
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	http.ServeFile(w, r, filePath)
}

func (h *BuildHandler) SyncBuild(build *store.Build) {
	if h.gh == nil || build.GHRunID == 0 {
		return
	}

	run, err := h.gh.GetWorkflowRun(build.GHRunID)
	if err != nil {
		return
	}

	newStatus := gh.RunStatusToBuildStatus(run.Status, run.Conclusion)
	if newStatus != build.Status {
		h.store.UpdateBuild(build.UUID, newStatus, build.Files, build.Message, build.GHRunID)
		build.Status = newStatus
	}

	if newStatus == "completed" {
		h.downloadArtifacts(build)
	}
}

func (h *BuildHandler) downloadArtifacts(build *store.Build) {
	if h.gh == nil || build.GHRunID == 0 {
		return
	}

	artifacts, err := h.gh.ListArtifacts(build.GHRunID)
	if err != nil {
		h.store.UpdateBuild(build.UUID, build.Status, build.Files, fmt.Sprintf("list artifacts: %v", err), build.GHRunID)
		return
	}

	dir := filepath.Join(h.storage, build.UUID)
	os.MkdirAll(dir, 0755)

	var downloadedFiles []string

	for _, artifact := range artifacts {
		zipPath := filepath.Join(dir, artifact.Name+".zip")
		f, err := os.Create(zipPath)
		if err != nil {
			continue
		}

		if err := h.gh.DownloadArtifact(artifact.ID, f); err != nil {
			f.Close()
			continue
		}
		f.Close()

		extractedFiles, err := extractZip(zipPath, dir)
		if err != nil {
			continue
		}
		os.Remove(zipPath)
		downloadedFiles = append(downloadedFiles, extractedFiles...)
	}

	if len(downloadedFiles) > 0 {
		filesJSON, _ := json.Marshal(downloadedFiles)
		h.store.UpdateBuild(build.UUID, build.Status, string(filesJSON), build.Message, build.GHRunID)
	}
}

func extractUUID(path, prefix string) string {
	s := strings.TrimPrefix(path, prefix)
	s = strings.Split(s, "/")[0]
	s = strings.Split(s, "?")[0]
	return strings.TrimSpace(s)
}

func extractZip(zipPath, destDir string) ([]string, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var files []string

	for _, f := range r.File {
		// Normalize zip entry path; reject path traversal
		name := filepath.ToSlash(f.Name)
		name = strings.TrimPrefix(name, "/")
		if name == "" || strings.HasPrefix(filepath.Base(name), ".") {
			continue
		}
		if strings.Contains(name, "..") {
			continue
		}

		target := filepath.Join(destDir, filepath.FromSlash(name))

		if f.FileInfo().IsDir() {
			os.MkdirAll(target, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			continue
		}

		out, err := os.Create(target)
		if err != nil {
			rc.Close()
			continue
		}

		_, copyErr := io.Copy(out, rc)
		out.Close()
		rc.Close()
		if copyErr != nil {
			continue
		}

		// Expose top-level files and the main portable exe for download list
		base := filepath.Base(name)
		if strings.EqualFold(filepath.Ext(base), ".exe") || !strings.Contains(name, "/") {
			files = append(files, base)
		} else if strings.HasSuffix(strings.ToLower(base), ".exe") {
			files = append(files, base)
		}
	}

	// If nested dirs were extracted, also list root-level entries for the panel
	if len(files) == 0 {
		entries, _ := os.ReadDir(destDir)
		for _, e := range entries {
			if !e.IsDir() {
				files = append(files, e.Name())
			}
		}
	}

	return files, nil
}
