package blinker

import (
	"context"
	"errors"
	"image/color"
	"sync"
	"testing"
	"time"
)

type fakeDevice struct {
	mu     sync.Mutex
	colors []color.Color
	err    error
}

func (f *fakeDevice) SetColor(c color.Color) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.colors = append(f.colors, c)
	return f.err
}

func (f *fakeDevice) seen() []color.Color {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]color.Color(nil), f.colors...)
}

var fast = Config{Interval: 5 * time.Millisecond, StartupPhase: 5 * time.Millisecond}

// runFor runs the blinker for d, then cancels it and waits for Run to return.
func runFor(t *testing.T, dev *fakeDevice, cfg Config, d time.Duration) []color.Color {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	done := make(chan struct{})
	go func() { Run(ctx, dev, cfg); close(done) }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return")
	}
	return dev.seen()
}

func TestBlinksBlueTwiceThenAlternates(t *testing.T) {
	got := runFor(t, &fakeDevice{}, fast, 100*time.Millisecond)
	want := []color.Color{Blue, Off, Blue, Off, Green, Red, Green, Red}
	if len(got) < len(want) {
		t.Fatalf("got %d changes, want at least %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i] != w {
			t.Fatalf("change %d = %v, want %v", i, got[i], w)
		}
	}
}

func TestSwitchesOffOnCancel(t *testing.T) {
	got := runFor(t, &fakeDevice{}, fast, 100*time.Millisecond)
	if last := got[len(got)-1]; last != Off {
		t.Fatalf("last color = %v, want off", last)
	}
}

func TestSwitchesOffWhenCancelledDuringStartup(t *testing.T) {
	cfg := Config{Interval: time.Hour, StartupPhase: time.Hour}
	got := runFor(t, &fakeDevice{}, cfg, 20*time.Millisecond)
	want := []color.Color{Blue, Off}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestStartupBlinkFitsInOneSecond(t *testing.T) {
	if 4*DefaultConfig.StartupPhase > time.Second {
		t.Fatalf("startup blink takes %v, want at most 1s", 4*DefaultConfig.StartupPhase)
	}
}

func TestContinuesAfterColorError(t *testing.T) {
	got := runFor(t, &fakeDevice{err: errors.New("boom")}, fast, 100*time.Millisecond)
	if len(got) < 7 {
		t.Fatalf("got %d changes, want the loop to continue past errors", len(got))
	}
}
