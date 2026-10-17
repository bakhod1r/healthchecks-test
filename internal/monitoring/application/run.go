// Package application wraps a job with healthchecks.io start/success/fail pings.
package application

import (
	"context"
	"log"

	"github.com/bakhod1r/healthchecks-test/internal/monitoring/domain"
)

// Pinger sends a signal for a check. Implemented in infrastructure.
type Pinger interface {
	Ping(ctx context.Context, c domain.Check, s domain.Signal, body string) error
}

// RunMonitored pings start, runs job, then pings success or fail.
// Ping errors are logged and never mask the job result.
func RunMonitored(ctx context.Context, p Pinger, c domain.Check, job func(context.Context) error) error {
	if err := p.Ping(ctx, c, domain.SignalStart, ""); err != nil {
		log.Printf("ping start: %v", err)
	}
	jobErr := job(ctx)
	sig, body := domain.SignalSuccess, ""
	if jobErr != nil {
		sig, body = domain.SignalFail, jobErr.Error()
	}
	if err := p.Ping(ctx, c, sig, body); err != nil {
		log.Printf("ping %s: %v", sig, err)
	}
	return jobErr
}
