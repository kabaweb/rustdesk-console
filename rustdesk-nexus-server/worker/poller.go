package worker

import (
	"log"
	"time"

	"rustdesk-nexus-server/store"
)

type BuildSyncer interface {
	SyncBuild(build *store.Build)
}

type Poller struct {
	syncer    BuildSyncer
	store     *store.Store
	interval  time.Duration
	stopCh    chan struct{}
}

func NewPoller(syncer BuildSyncer, s *store.Store, interval time.Duration) *Poller {
	return &Poller{
		syncer:   syncer,
		store:    s,
		interval: interval,
		stopCh:   make(chan struct{}),
	}
}

func (p *Poller) Start() {
	go func() {
		ticker := time.NewTicker(p.interval)
		defer ticker.Stop()
		p.poll()
		for {
			select {
			case <-ticker.C:
				p.poll()
			case <-p.stopCh:
				return
			}
		}
	}()
	log.Printf("Worker poller started (interval: %s)", p.interval)
}

func (p *Poller) Stop() {
	close(p.stopCh)
}

func (p *Poller) poll() {
	builds, err := p.store.GetActiveBuilds()
	if err != nil {
		log.Printf("Poller: error getting active builds: %v", err)
		return
	}

	if len(builds) > 0 {
		log.Printf("Poller: syncing %d active builds", len(builds))
	}

	for _, build := range builds {
		p.syncer.SyncBuild(build)
	}
}
