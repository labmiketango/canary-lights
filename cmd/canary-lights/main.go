// Command canary-lights drives a USB LED to show pipeline status.
// For now it alternates the LED between green and red every second.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/boombuler/led"

	"github.com/labmiketango/canary-lights/internal/blinker"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var dev led.Device
	for info := range led.Devices() {
		d, err := info.Open()
		if err != nil {
			log.Printf("open %s: %v", info.GetPath(), err)
			continue
		}
		dev = d
		break
	}
	if dev == nil {
		log.Fatal("no usable LED device found")
	}
	defer dev.Close() // closing turns the LED off

	blinker.Run(ctx, dev, time.Second)
}
