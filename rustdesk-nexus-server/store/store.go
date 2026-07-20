package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Build struct {
	UUID      string    `json:"uuid"`
	UserGUID  string    `json:"user_guid"`
	OS        string    `json:"os"`
	Arch      string    `json:"arch"`
	AppName   string    `json:"app_name"`
	Custom    string    `json:"custom"`
	Status    string    `json:"status"`
	Files     string    `json:"files"`
	Message   string    `json:"message"`
	GHRunID   int64     `json:"gh_run_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	UserGUID      string    `json:"user_guid"`
	Token         string    `json:"token"`
	Username      string    `json:"username"`
	ExpiresAt     time.Time `json:"expires_at"`
	CurrentUUID   string    `json:"current_uuid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Store struct {
	db *sql.DB
	mu sync.RWMutex
}

func New(path string) (*Store, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("create data dir: %w", err)
	}

	db, err := sql.Open("sqlite3", path+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	db.SetMaxOpenConns(1)

	s := &Store{db: db}
	if err := s.migrate(); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return s, nil
}

func (s *Store) migrate() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS tokens (
			user_guid TEXT PRIMARY KEY,
			token TEXT NOT NULL,
			username TEXT NOT NULL DEFAULT '',
			expires_at DATETIME NOT NULL,
			current_uuid TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS builds (
			uuid TEXT PRIMARY KEY,
			user_guid TEXT NOT NULL,
			os TEXT NOT NULL DEFAULT 'windows',
			arch TEXT NOT NULL DEFAULT 'x86_64',
			app_name TEXT NOT NULL DEFAULT '',
			custom TEXT NOT NULL DEFAULT '{}',
			status TEXT NOT NULL DEFAULT 'pending',
			files TEXT NOT NULL DEFAULT '[]',
			message TEXT NOT NULL DEFAULT '',
			gh_run_id INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_builds_user_guid ON builds(user_guid)`,
		`CREATE INDEX IF NOT EXISTS idx_builds_status ON builds(status)`,
		`CREATE INDEX IF NOT EXISTS idx_builds_created ON builds(created_at)`,
	}

	for _, q := range queries {
		if _, err := s.db.Exec(q); err != nil {
			return fmt.Errorf("exec %q: %w", q, err)
		}
	}
	return nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

// ── Token operations ──

func (s *Store) UpsertToken(t *Token) error {
	_, err := s.db.Exec(
		`INSERT INTO tokens (user_guid, token, username, expires_at, current_uuid) 
		 VALUES (?, ?, ?, ?, ?)
		 ON CONFLICT(user_guid) DO UPDATE SET 
		 token=excluded.token, username=excluded.username, 
		 expires_at=excluded.expires_at, current_uuid=excluded.current_uuid,
		 updated_at=CURRENT_TIMESTAMP`,
		t.UserGUID, t.Token, t.Username, t.ExpiresAt, t.CurrentUUID,
	)
	return err
}

func (s *Store) GetToken(userGUID string) (*Token, error) {
	t := &Token{}
	err := s.db.QueryRow(
		`SELECT user_guid, token, username, expires_at, COALESCE(current_uuid,''), created_at, updated_at 
		 FROM tokens WHERE user_guid=?`, userGUID,
	).Scan(&t.UserGUID, &t.Token, &t.Username, &t.ExpiresAt, &t.CurrentUUID, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return t, err
}

func (s *Store) DeleteToken(userGUID string) error {
	_, err := s.db.Exec(`DELETE FROM tokens WHERE user_guid=?`, userGUID)
	return err
}

// ── Build operations ──

func (s *Store) CreateBuild(b *Build) error {
	_, err := s.db.Exec(
		`INSERT INTO builds (uuid, user_guid, os, arch, app_name, custom, status, files, message, gh_run_id)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		b.UUID, b.UserGUID, b.OS, b.Arch, b.AppName, b.Custom, b.Status, b.Files, b.Message, b.GHRunID,
	)
	return err
}

func (s *Store) UpdateBuild(uuid string, status string, files string, message string, ghRunID int64) error {
	_, err := s.db.Exec(
		`UPDATE builds SET status=?, files=?, message=?, gh_run_id=?, updated_at=CURRENT_TIMESTAMP WHERE uuid=?`,
		status, files, message, ghRunID, uuid,
	)
	return err
}

func (s *Store) GetBuild(uuid string) (*Build, error) {
	b := &Build{}
	err := s.db.QueryRow(
		`SELECT uuid, user_guid, os, arch, app_name, custom, status, files, message, gh_run_id, created_at, updated_at
		 FROM builds WHERE uuid=?`, uuid,
	).Scan(&b.UUID, &b.UserGUID, &b.OS, &b.Arch, &b.AppName, &b.Custom, &b.Status, &b.Files, &b.Message, &b.GHRunID, &b.CreatedAt, &b.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return b, err
}

func (s *Store) ListBuilds(userGUID string) ([]*Build, error) {
	rows, err := s.db.Query(
		`SELECT uuid, user_guid, os, arch, app_name, custom, status, files, message, gh_run_id, created_at, updated_at
		 FROM builds WHERE user_guid=? ORDER BY created_at DESC LIMIT 50`, userGUID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var builds []*Build
	for rows.Next() {
		b := &Build{}
		if err := rows.Scan(&b.UUID, &b.UserGUID, &b.OS, &b.Arch, &b.AppName, &b.Custom, &b.Status, &b.Files, &b.Message, &b.GHRunID, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		builds = append(builds, b)
	}
	return builds, rows.Err()
}

func (s *Store) FindActiveBuild(userGUID string) (*Build, error) {
	b := &Build{}
	err := s.db.QueryRow(
		`SELECT uuid, user_guid, os, arch, app_name, custom, status, files, message, gh_run_id, created_at, updated_at
		 FROM builds WHERE user_guid=? AND status IN ('pending','building') LIMIT 1`, userGUID,
	).Scan(&b.UUID, &b.UserGUID, &b.OS, &b.Arch, &b.AppName, &b.Custom, &b.Status, &b.Files, &b.Message, &b.GHRunID, &b.CreatedAt, &b.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return b, err
}

func (s *Store) CountBuildsThisMonth(userGUID string) (int, error) {
	var count int
	err := s.db.QueryRow(
		`SELECT COUNT(*) FROM builds WHERE user_guid=? AND created_at >= date('now','start of month')`,
		userGUID,
	).Scan(&count)
	return count, err
}

func (s *Store) DeleteBuild(uuid string) error {
	_, err := s.db.Exec(`DELETE FROM builds WHERE uuid=?`, uuid)
	return err
}

func (s *Store) GetActiveBuilds() ([]*Build, error) {
	rows, err := s.db.Query(
		`SELECT uuid, user_guid, os, arch, app_name, custom, status, files, message, gh_run_id, created_at, updated_at
		 FROM builds WHERE status IN ('pending','building')`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var builds []*Build
	for rows.Next() {
		b := &Build{}
		if err := rows.Scan(&b.UUID, &b.UserGUID, &b.OS, &b.Arch, &b.AppName, &b.Custom, &b.Status, &b.Files, &b.Message, &b.GHRunID, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		builds = append(builds, b)
	}
	return builds, rows.Err()
}
