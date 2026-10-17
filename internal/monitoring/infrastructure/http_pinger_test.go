package infrastructure

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakhod1r/healthchecks-test/internal/monitoring/domain"
)

func TestHTTPPingerPostsToSignalPath(t *testing.T) {
	var gotPath, gotBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
	}))
	defer srv.Close()

	c, _ := domain.NewCheck(srv.URL + "/ping/abc")
	if err := NewHTTPPinger().Ping(context.Background(), c, domain.SignalFail, "boom"); err != nil {
		t.Fatal(err)
	}
	if gotPath != "/ping/abc/fail" || gotBody != "boom" {
		t.Fatalf("got path=%q body=%q", gotPath, gotBody)
	}
}

func TestHTTPPingerNon200IsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	c, _ := domain.NewCheck(srv.URL + "/ping/missing")
	if err := NewHTTPPinger().Ping(context.Background(), c, domain.SignalSuccess, ""); err == nil {
		t.Fatal("want error on 404")
	}
}
