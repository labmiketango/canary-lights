// Package blinker alternates an LED between green and red.
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
)

// Device is the subset of an LED device the blinker needs.
type Device interface {
	SetColor(c color.Color) error
}

// Run sets the LED to green, then alternates green/red every interval until
// ctx is cancelled. Color errors are logged and do not stop the loop.
func Run(ctx context.Context, dev Device, interval time.Duration) {
	colors := []color.Color{Green, Red}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for i := 0; ; i++ {
		if err := dev.SetColor(colors[i%len(colors)]); err != nil {
			log.Printf("set color: %v", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
