// Package infrastructure contains the HTTP adapter for healthchecks.io pings.
package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bakhod1r/healthchecks-test/internal/monitoring/domain"
)

type HTTPPinger struct {
	Client *http.Client
}

func NewHTTPPinger() *HTTPPinger {
	return &HTTPPinger{Client: &http.Client{Timeout: 10 * time.Second}}
}

func (p *HTTPPinger) Ping(ctx context.Context, c domain.Check, s domain.Signal, body string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URLFor(s), strings.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := p.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ping %s: unexpected status %d", s, resp.StatusCode)
	}
	return nil
}
