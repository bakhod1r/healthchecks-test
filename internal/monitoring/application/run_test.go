package application

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/bakhod1r/healthchecks-test/internal/monitoring/domain"
)

type fakePinger struct {
	signals []domain.Signal
	bodies  []string
	err     error
}

func (f *fakePinger) Ping(_ context.Context, _ domain.Check, s domain.Signal, body string) error {
	f.signals = append(f.signals, s)
	f.bodies = append(f.bodies, body)
	return f.err
}

func check(t *testing.T) domain.Check {
	c, err := domain.NewCheck("http://hc/ping/x")
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestRunMonitoredSuccess(t *testing.T) {
	p := &fakePinger{}
	if err := RunMonitored(context.Background(), p, check(t), func(context.Context) error { return nil }); err != nil {
		t.Fatal(err)
	}
	want := []domain.Signal{domain.SignalStart, domain.SignalSuccess}
	if !reflect.DeepEqual(p.signals, want) {
		t.Fatalf("got %v want %v", p.signals, want)
	}
}

func TestRunMonitoredFailSendsErrorBody(t *testing.T) {
	p := &fakePinger{}
	boom := errors.New("boom")
	err := RunMonitored(context.Background(), p, check(t), func(context.Context) error { return boom })
	if !errors.Is(err, boom) {
		t.Fatalf("want boom, got %v", err)
	}
	if p.signals[1] != domain.SignalFail || p.bodies[1] != "boom" {
		t.Fatalf("got %v %v", p.signals, p.bodies)
	}
}

func TestRunMonitoredPingErrorDoesNotMaskJob(t *testing.T) {
	p := &fakePinger{err: errors.New("network")}
	if err := RunMonitored(context.Background(), p, check(t), func(context.Context) error { return nil }); err != nil {
		t.Fatalf("want nil, got %v", err)
	}
}
