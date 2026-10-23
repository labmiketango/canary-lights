// Package blinker alternates an LED between green and red.
package blinker

import (
	"context"
	"image/color"
	"log"
	"sync"
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
// ctx is cancelled. The timing never waits for the device: colors are sent by
// a separate goroutine, and if the device is slow only the newest pending
// color is kept. Color errors are logged and do not stop the loop. Run returns
// once any in-flight device call has finished.
func Run(ctx context.Context, dev Device, interval time.Duration) {
	pending := make(chan color.Color, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range pending {
			if err := dev.SetColor(c); err != nil {
				log.Printf("set color: %v", err)
			}
		}
	}()
	defer func() {
		close(pending)
		wg.Wait()
	}()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	colors := []color.Color{Green, Red}
	for i := 0; ; i++ {
		send(pending, colors[i%len(colors)])
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

// send queues c, replacing a queued color the worker has not picked up yet.
func send(pending chan color.Color, c color.Color) {
	for {
		select {
		case pending <- c:
			return
		default:
		}
		select {
		case <-pending:
		default:
		}
	}
}
