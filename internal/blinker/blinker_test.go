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
	got    chan struct{}
}

func (f *fakeDevice) SetColor(c color.Color) error {
	f.mu.Lock()
	f.colors = append(f.colors, c)
	f.mu.Unlock()
	f.got <- struct{}{}
	return f.err
}

func runUntil(t *testing.T, dev *fakeDevice, n int) []color.Color {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() { Run(ctx, dev, 5*time.Millisecond); close(done) }()
	for i := 0; i < n; i++ {
		select {
		case <-dev.got:
		case <-time.After(2 * time.Second):
			t.Fatal("timed out waiting for color change")
		}
	}
	cancel()
	for {
		select {
		case <-dev.got: // drain
		case <-done:
			dev.mu.Lock()
			defer dev.mu.Unlock()
			return append([]color.Color(nil), dev.colors...)
		}
	}
}

func TestAlternatesGreenAndRed(t *testing.T) {
	got := runUntil(t, &fakeDevice{got: make(chan struct{}, 100)}, 4)
	want := []color.Color{Green, Red, Green, Red}
	for i, w := range want {
		if got[i] != w {
			t.Fatalf("change %d = %v, want %v", i, got[i], w)
		}
	}
}

func TestContinuesAfterColorError(t *testing.T) {
	got := runUntil(t, &fakeDevice{got: make(chan struct{}, 100), err: errors.New("boom")}, 3)
	if len(got) < 3 {
		t.Fatalf("got %d changes, want at least 3", len(got))
	}
}

func TestStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	dev := &fakeDevice{got: make(chan struct{}, 100)}
	done := make(chan struct{})
	go func() { Run(ctx, dev, time.Hour); close(done) }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not return after cancel")
	}
}
