# Audit History

Append-only. One row per change.

| Date       | Spec          | Branch               | Change |
|------------|---------------|----------------------|--------|
| 2026-10-23 | specs/init.md | feature/init-project | Initial Go project structure: `cmd/canary-lights`, `internal/blinker`, unit tests. LED alternates green/red every second. Dependency `github.com/boombuler/led` (upstream of go-busylight, which does not compile on macOS with Go 1.26). |
| 2026-10-23 | specs/init.md | feature/init-project | Added `.vscode/launch.json` (run, debug, test); added `bin/` to `.gitignore`. Replaced earlier TypeScript documentation with Go documentation. |
| 2026-10-23 | specs/init.md | feature/init-project | Blinker no longer waits for the device: colors are sent from a separate goroutine that keeps only the newest pending color, so a slow or "not responding" device does not stretch the 1s interval. Added test. |
