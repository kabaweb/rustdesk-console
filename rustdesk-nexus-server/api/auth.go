package api

import (
	"fmt"
	"net/http"
	"time"

	"rustdesk-nexus-server/config"
	"rustdesk-nexus-server/store"

	"github.com/google/uuid"
)

type AuthHandler struct {
	cfg   *config.Config
	store *store.Store
}

func NewAuthHandler(cfg *config.Config, s *store.Store) *AuthHandler {
	return &AuthHandler{cfg: cfg, store: s}
}

type loginSession struct {
	loginID   string
	expiresAt time.Time
	completed bool
	token     string
	username  string
}

var sessions = map[string]*loginSession{}

func (h *AuthHandler) GitHubLogin(w http.ResponseWriter, r *http.Request) {
	sessionID := uuid.New().String()
	expiresIn := 600

	sessions[sessionID] = &loginSession{
		loginID:   sessionID,
		expiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second),
	}

	callbackURL := h.cfg.PublicURL + "/v1/auth/github/callback?login_id=" + sessionID
	if h.cfg.PublicURL == "" {
		callbackURL = fmt.Sprintf("http://localhost:%d/v1/auth/github/callback?login_id=%s", h.cfg.Port, sessionID)
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"login_id":   sessionID,
		"auth_url":   callbackURL,
		"expires_in": expiresIn,
	})
}

func (h *AuthHandler) GitHubCallback(w http.ResponseWriter, r *http.Request) {
	loginID := r.URL.Query().Get("login_id")
	session, ok := sessions[loginID]
	if !ok || time.Now().After(session.expiresAt) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<html><body><h2>Session expired</h2><p>Please go back and start again.</p></body></html>`))
		return
	}

	token, err := GenerateToken(h.cfg.JWTSecret, "default", "admin", 2592000)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "token generation failed"})
		return
	}

	h.store.UpsertToken(&store.Token{
		UserGUID:  "default",
		Token:     token,
		Username:  "admin",
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	})

	session.completed = true
	session.token = token
	session.username = "admin"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<html><body><h2>Authorization successful!</h2><p>You can close this page and return to the console.</p></body></html>`))
}

func (h *AuthHandler) GitHubStatus(w http.ResponseWriter, r *http.Request) {
	loginID := r.URL.Query().Get("login_id")
	if loginID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing login_id"})
		return
	}

	session, ok := sessions[loginID]
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "session not found", "state": "failed"})
		return
	}

	if time.Now().After(session.expiresAt) {
		delete(sessions, loginID)
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"state": "failed",
			"error": "session expired",
		})
		return
	}

	if session.completed {
		delete(sessions, loginID)
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"state":      "completed",
			"token":      session.token,
			"username":   "admin",
			"expires_in": 2592000,
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"state": "pending"})
}

func (h *AuthHandler) BindStatus(w http.ResponseWriter, r *http.Request) {
	userGUID := r.Header.Get("X-User-GUID")
	t, _ := h.store.GetToken(userGUID)
	if t != nil && t.ExpiresAt.Before(time.Now()) {
		writeJSON(w, http.StatusOK, map[string]interface{}{"bound": false, "expired": true, "nexus_username": t.Username})
		return
	}
	if t != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{"bound": true, "nexus_username": t.Username})
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"bound": false})
}

func (h *AuthHandler) Unbind(w http.ResponseWriter, r *http.Request) {
	userGUID := r.Header.Get("X-User-GUID")
	h.store.DeleteToken(userGUID)
	writeJSON(w, http.StatusOK, map[string]string{"message": "unbind ok"})
}
