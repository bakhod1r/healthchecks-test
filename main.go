// Minimal backend for the healthchecks.io test: GET /health returns {"status":"ok"}.
// If HC_PING_URL is set, it also pings healthchecks every HC_PING_INTERVAL (default 1m).
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	http.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}` + "\n"))
	})

	if url := os.Getenv("HC_PING_URL"); url != "" {
		interval, err := time.ParseDuration(os.Getenv("HC_PING_INTERVAL"))
		if err != nil {
			interval = time.Minute
		}
		go heartbeat(url, interval)
	}

	log.Printf("backend listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// heartbeat pings healthchecks while the process is alive; if it dies, the check goes down.
func heartbeat(url string, interval time.Duration) {
	client := &http.Client{Timeout: 10 * time.Second}
	log.Printf("heartbeat: pinging %s every %s", url, interval)
	for {
		log.Printf("heartbeat: sending ping")
		start := time.Now()
		resp, err := client.Get(url)
		if err != nil {
			log.Printf("heartbeat: failed: %v", err)
		} else {
			resp.Body.Close()
			log.Printf("heartbeat: sent, status %d in %s", resp.StatusCode, time.Since(start).Round(time.Millisecond))
		}
		time.Sleep(interval)
	}
}
