package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             int
	JWTSecret        string
	GitHubToken      string
	GitHubOwner      string
	GitHubRepo       string
	WorkflowPrefix   string
	RustDeskRepo     string
	RustDeskRef      string
	StoragePath      string
	MaxBuildsPerMonth int
	PublicURL        string
}

func Load() *Config {
	return &Config{
		Port:              getEnvInt("NEXUS_PORT", 8090),
		JWTSecret:         getEnv("NEXUS_JWT_SECRET", "change-me"),
		GitHubToken:       getEnv("GITHUB_TOKEN", ""),
		GitHubOwner:       getEnv("GITHUB_OWNER", ""),
		GitHubRepo:        getEnv("GITHUB_REPO", "rustdesk-generator"),
		WorkflowPrefix:    getEnv("GITHUB_WORKFLOW_PREFIX", "build-custom"),
		RustDeskRepo:      getEnv("RUSTDESK_REPO", "rustdesk/rustdesk"),
		RustDeskRef:       getEnv("RUSTDESK_REF", "refs/tags/1.4.9"),
		StoragePath:       getEnv("NEXUS_STORAGE_PATH", "./artifacts"),
		MaxBuildsPerMonth: getEnvInt("NEXUS_MAX_BUILDS_PER_MONTH", 15),
		PublicURL:         getEnv("NEXUS_PUBLIC_URL", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return defaultVal
}
