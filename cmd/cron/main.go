// Command cron runs a job every CRON_INTERVAL and reports it to healthchecks.io.
// Job: GET BACKEND_URL/health must return 200.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bakhod1r/healthchecks-test/internal/monitoring/application"
	"github.com/bakhod1r/healthchecks-test/internal/monitoring/domain"
	"github.com/bakhod1r/healthchecks-test/internal/monitoring/infrastructure"
)

func main() {
	check, err := domain.NewCheck(os.Getenv("HC_PING_URL"))
	if err != nil {
		log.Fatalf("HC_PING_URL: %v", err)
	}
	backendURL := envOr("BACKEND_URL", "http://localhost:8080")
	interval, err := time.ParseDuration(envOr("CRON_INTERVAL", "1m"))
	if err != nil {
		log.Fatalf("CRON_INTERVAL: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pinger := infrastructure.NewHTTPPinger()
	client := &http.Client{Timeout: 5 * time.Second}
	job := func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, backendURL+"/health", nil)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("backend health: status %d", resp.StatusCode)
		}
		return nil
	}

	run := func() {
		if err := application.RunMonitored(ctx, pinger, check, job); err != nil {
			log.Printf("job failed: %v", err)
			return
		}
		log.Print("job ok")
	}

	run()
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			run()
		}
	}
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
