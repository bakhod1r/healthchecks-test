// Minimal backend for the healthchecks.io cronjob test: GET /health returns {"status":"ok"}.
package main

import (
	"log"
	"net/http"
	"os"
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
	log.Printf("backend listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
