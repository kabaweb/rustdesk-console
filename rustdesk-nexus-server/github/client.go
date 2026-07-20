package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	token     string
	owner     string
	repo      string
	workflow  string
	http      *http.Client
}

func New(token, owner, repo, workflow string) *Client {
	return &Client{
		token:    token,
		owner:    owner,
		repo:     repo,
		workflow: workflow,
		http:     &http.Client{Timeout: 30 * time.Second},
	}
}

type WorkflowDispatchInput struct {
	RequestID    string `json:"request_id"`
	RustDeskRepo string `json:"rustdesk_repo"`
	RustDeskRef  string `json:"rustdesk_ref"`
	AppName      string `json:"app_name"`
	CustomConfig string `json:"custom_config"`
	OS           string `json:"os"`
	Arch         string `json:"arch"`
}

type WorkflowRun struct {
	ID         int64  `json:"id"`
	Status     string `json:"status"`
	Conclusion string `json:"conclusion"`
	HTMLURL    string `json:"html_url"`
}

type WorkflowRunsResponse struct {
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
}

type Artifact struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	SizeInBytes        int64  `json:"size_in_bytes"`
	WorkflowRun        struct{ ID int64 } `json:"workflow_run"`
	ArchiveDownloadURL string `json:"archive_download_url"`
}

type ArtifactsResponse struct {
	Artifacts []Artifact `json:"artifacts"`
}

func (c *Client) DispatchWorkflow(input *WorkflowDispatchInput) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%s/dispatches",
		c.owner, c.repo, c.workflow)

	body := map[string]interface{}{
		"ref": "main",
		"inputs": map[string]interface{}{
			"request_id":    input.RequestID,
			"rustdesk_repo": input.RustDeskRepo,
			"rustdesk_ref":  input.RustDeskRef,
			"app_name":      input.AppName,
			"custom_config": input.CustomConfig,
			"os":            input.OS,
			"arch":          input.Arch,
		},
	}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("dispatch workflow: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("dispatch workflow: status %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

func (c *Client) GetLatestRunForRequest(requestID string) (*WorkflowRun, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%s/runs?per_page=20",
		c.owner, c.repo, c.workflow)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get runs: %w", err)
	}
	defer resp.Body.Close()

	var runsResp WorkflowRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&runsResp); err != nil {
		return nil, fmt.Errorf("decode runs: %w", err)
	}

	for _, run := range runsResp.WorkflowRuns {
		return &run, nil
	}
	return nil, nil
}

func (c *Client) GetWorkflowRun(runID int64) (*WorkflowRun, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs/%d",
		c.owner, c.repo, runID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get run: %w", err)
	}
	defer resp.Body.Close()

	var run WorkflowRun
	if err := json.NewDecoder(resp.Body).Decode(&run); err != nil {
		return nil, fmt.Errorf("decode run: %w", err)
	}
	return &run, nil
}

func (c *Client) ListArtifacts(runID int64) ([]Artifact, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs/%d/artifacts",
		c.owner, c.repo, runID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("list artifacts: %w", err)
	}
	defer resp.Body.Close()

	var artifactsResp ArtifactsResponse
	if err := json.NewDecoder(resp.Body).Decode(&artifactsResp); err != nil {
		return nil, fmt.Errorf("decode artifacts: %w", err)
	}
	return artifactsResp.Artifacts, nil
}

func (c *Client) DownloadArtifact(artifactID int64, writer io.Writer) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/artifacts/%d/zip",
		c.owner, c.repo, artifactID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("download artifact: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 || resp.StatusCode == 307 {
		loc := resp.Header.Get("Location")
		if loc == "" {
			return fmt.Errorf("redirect without Location header")
		}
		req2, _ := http.NewRequest("GET", loc, nil)
		resp2, err := c.http.Do(req2)
		if err != nil {
			return fmt.Errorf("follow redirect: %w", err)
		}
		defer resp2.Body.Close()
		_, err = io.Copy(writer, resp2.Body)
		return err
	}

	_, err = io.Copy(writer, resp.Body)
	return err
}

func (c *Client) CancelWorkflowRun(runID int64) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs/%d/cancel",
		c.owner, c.repo, runID)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("cancel run: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		return fmt.Errorf("cancel run: status %d", resp.StatusCode)
	}
	return nil
}

func RunStatusToBuildStatus(status, conclusion string) string {
	switch status {
	case "queued", "in_progress":
		return "building"
	case "completed":
		switch conclusion {
		case "success":
			return "completed"
		case "failure", "cancelled", "skipped":
			return conclusion
		default:
			return "failed"
		}
	default:
		return "pending"
	}
}

func (c *Client) GetRunIDFromDispatchResponse(respBody []byte) (int64, error) {
	var result struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return 0, err
	}
	return result.ID, nil
}

func ParseInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}
