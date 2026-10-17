package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bakhod1r/healthchecks-test/internal/backend"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("backend listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, backend.NewHandler()))
}
