# Audit History

Append-only. One row per change.

| Date       | Spec          | Branch               | Change |
|------------|---------------|----------------------|--------|
| 2026-10-23 | specs/init.md | feature/init-project | Initial Go project structure: `cmd/canary-lights`, `internal/blinker`, unit tests. LED alternates green/red every second. Dependency `github.com/boombuler/led` (upstream of go-busylight, which does not compile on macOS with Go 1.26). |
| 2026-10-23 | specs/init.md | feature/init-project | Added `.vscode/launch.json` (run, debug, test); added `bin/` to `.gitignore`. Replaced earlier TypeScript documentation with Go documentation. |
| 2026-10-23 | specs/init.md | feature/init-project | Spec update: LED blinks blue twice within 1s on start (4 x 250ms phases), then alternates green/red every second; LED is switched off explicitly when the app exits or terminates, including during the startup blink. `blinker.Run` now takes a `Config`. Tests updated. |
| 2026-10-23 | specs/init.md | feature/init-project | Release: committed updated specs/init.md and CLAUDE.MD; merged feature/init-project into main. Verified on hardware (BlinkStick). |
