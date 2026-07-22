package api

import (
	"net/http"
	"strings"

	"rustdesk-nexus-server/config"
	"rustdesk-nexus-server/github"
	"rustdesk-nexus-server/store"
)

func NewRouter(cfg *config.Config, s *store.Store) (http.Handler, *BuildHandler) {
	var ghClient *github.Client
	if cfg.GitHubToken != "" && cfg.GitHubOwner != "" && cfg.GitHubRepo != "" {
		ghClient = github.New(cfg.GitHubToken, cfg.GitHubOwner, cfg.GitHubRepo, cfg.WorkflowPrefix)
	}

	authHandler := NewAuthHandler(cfg, s)
	buildHandler := NewBuildHandler(cfg, s, ghClient, cfg.StoragePath)

	mux := http.NewServeMux()
	jwtAuth := JWTMiddleware(cfg.JWTSecret)

	mux.HandleFunc("/v1/auth/github/login", authHandler.GitHubLogin)
	mux.HandleFunc("/v1/auth/github/status", authHandler.GitHubStatus)
	mux.HandleFunc("/v1/auth/github/callback", authHandler.GitHubCallback)

	mux.HandleFunc("/v1/client/generate", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			jwtAuth(http.HandlerFunc(buildHandler.Generate)).ServeHTTP(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	mux.HandleFunc("/v1/client/generate/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasPrefix(path, "/v1/client/generate/") {
			parts := strings.SplitN(strings.TrimPrefix(path, "/v1/client/generate/"), "/", 2)
			if len(parts) == 2 && parts[1] != "" {
				jwtAuth(http.HandlerFunc(buildHandler.Download)).ServeHTTP(w, r)
				return
			}
			jwtAuth(http.HandlerFunc(buildHandler.GetStatus)).ServeHTTP(w, r)
			return
		}
	})

	mux.HandleFunc("/v1/client/download/", func(w http.ResponseWriter, r *http.Request) {
		jwtAuth(http.HandlerFunc(buildHandler.Download)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	return mux, buildHandler
}
