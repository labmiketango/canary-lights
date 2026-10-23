// Package blinker drives an LED: two blue blinks on start, then green/red
// alternation until stopped, then off.
package blinker

import (
	"context"
	"image/color"
	"log"
	"time"
)

var (
	Green = color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	Red   = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	Blue  = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}
	Off   = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
)

// Device is the subset of an LED device the blinker needs.
type Device interface {
	SetColor(c color.Color) error
}

// Config holds the timings.
type Config struct {
	// Interval between green/red changes.
	Interval time.Duration
	// StartupPhase is the length of each on/off phase of the startup blink.
	// Two blue blinks are four phases (on, off, on, off).
	StartupPhase time.Duration
}

// DefaultConfig blinks blue twice within one second, then switches every second.
var DefaultConfig = Config{Interval: time.Second, StartupPhase: 250 * time.Millisecond}

// Run blinks blue twice, then alternates green/red every Interval until ctx is
// cancelled. The LED is switched off before Run returns, also when ctx is
// cancelled during the startup blink. Color errors are logged and do not stop
// the loop.
func Run(ctx context.Context, dev Device, cfg Config) {
	defer set(dev, Off)

	for i := 0; i < 4; i++ {
		c := color.Color(Off)
		if i%2 == 0 {
			c = Blue
		}
		set(dev, c)
		if !sleep(ctx, cfg.StartupPhase) {
			return
		}
	}

	colors := []color.Color{Green, Red}
	for i := 0; ; i++ {
		set(dev, colors[i%len(colors)])
		if !sleep(ctx, cfg.Interval) {
			return
		}
	}
}

func set(dev Device, c color.Color) {
	if err := dev.SetColor(c); err != nil {
		log.Printf("set color: %v", err)
	}
}

// sleep waits for d and reports false if ctx was cancelled first.
func sleep(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-t.C:
		return true
	}
}
