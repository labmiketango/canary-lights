# Implementation Summary: Project Init

## Structure

- `cmd/canary-lights/main.go`: opens the first usable LED device, handles SIGINT/SIGTERM, closes the device on exit (which turns the LED off).
- `internal/blinker`: `Run(ctx, dev, interval)` sets green, then alternates green/red each interval until the context is cancelled. Colors are sent from a separate goroutine (only the newest pending color is kept), so a slow device does not delay the 1s timing. Color errors are logged and do not stop the loop. It depends on a one-method `Device` interface, so it is tested with a fake.
- `.vscode/launch.json`: Run, Debug and Test configurations.

## Dependency note

The spec names `BaaaZen/go-busylight`. Its `go-hid` dependency does not compile on macOS with Go 1.26 (CoreFoundation nil-type errors). Dependencies must not be modified, so this repo uses `github.com/boombuler/led`, the upstream of go-busylight with the same API. Behavior with a physical device is untested.

## Verification (Go 1.26.4, macOS arm64)

Unit tests with `-race`: pass. `go vet`: clean. `gofmt`: clean. Build (type check): OK. `go mod verify`: OK. Dependency analysis: `go list -m -u all` shows no updates. No vulnerability scanner was used (not part of the chosen tooling).
