package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"rustdesk-nexus-server/api"
	"rustdesk-nexus-server/config"
	"rustdesk-nexus-server/store"
	"rustdesk-nexus-server/worker"
)

func main() {
	cfg := config.Load()

	s, err := store.New("./data/nexus.db")
	if err != nil {
		log.Fatalf("Failed to init store: %v", err)
	}
	defer s.Close()

	if err := os.MkdirAll(cfg.StoragePath, 0755); err != nil {
		log.Fatalf("Failed to create storage dir: %v", err)
	}

	router, buildHandler := api.NewRouter(cfg, s)

	poller := worker.NewPoller(buildHandler, s, 15*time.Second)
	poller.Start()
	defer poller.Stop()

	bindAddr := "0.0.0.0"
	if h := os.Getenv("NEXUS_HOST"); h != "" {
		bindAddr = h
	}
	bindAddr = fmt.Sprintf("%s:%d", bindAddr, cfg.Port)

	log.Printf("Nexus API server starting on %s", bindAddr)
	log.Printf("  Storage: %s", cfg.StoragePath)
	log.Printf("  GitHub: %s/%s (workflow prefix: %s)", cfg.GitHubOwner, cfg.GitHubRepo, cfg.WorkflowPrefix)
	log.Printf("  RustDesk source: %s @ %s", cfg.RustDeskRepo, cfg.RustDeskRef)

	if err := http.ListenAndServe(bindAddr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
