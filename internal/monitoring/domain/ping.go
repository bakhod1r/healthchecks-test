// Package domain holds pure healthchecks.io ping rules.
package domain

import (
	"errors"
	"strings"
)

// Signal is the kind of ping sent to healthchecks.io.
type Signal string

const (
	SignalStart   Signal = "start"
	SignalSuccess Signal = "success"
	SignalFail    Signal = "fail"
)

var ErrEmptyPingURL = errors.New("ping url is empty")

// Check is a monitored job identified by its ping URL (e.g. https://hc-ping.com/<uuid>).
type Check struct {
	pingURL string
}

func NewCheck(pingURL string) (Check, error) {
	u := strings.TrimRight(strings.TrimSpace(pingURL), "/")
	if u == "" {
		return Check{}, ErrEmptyPingURL
	}
	return Check{pingURL: u}, nil
}

// URLFor returns the endpoint for a signal: success is the bare URL, others get a suffix.
func (c Check) URLFor(s Signal) string {
	if s == SignalSuccess {
		return c.pingURL
	}
	return c.pingURL + "/" + string(s)
}
