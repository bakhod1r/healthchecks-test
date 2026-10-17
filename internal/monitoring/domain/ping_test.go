package domain

import "testing"

func TestNewCheckRejectsEmpty(t *testing.T) {
	if _, err := NewCheck("  "); err != ErrEmptyPingURL {
		t.Fatalf("want ErrEmptyPingURL, got %v", err)
	}
}

func TestURLFor(t *testing.T) {
	c, err := NewCheck("https://hc-ping.com/abc/")
	if err != nil {
		t.Fatal(err)
	}
	cases := map[Signal]string{
		SignalSuccess: "https://hc-ping.com/abc",
		SignalStart:   "https://hc-ping.com/abc/start",
		SignalFail:    "https://hc-ping.com/abc/fail",
	}
	for s, want := range cases {
		if got := c.URLFor(s); got != want {
			t.Errorf("%s: got %q want %q", s, got, want)
		}
	}
}
