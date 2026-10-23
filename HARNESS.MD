# Harness Pipeline Instructions

## General

Use these instructions to build a pipeline within Harness. Note that Standing Orders and OPA (Open Policy Agent) policies have already been defined by the Platform Team. These policies must be adhered to and are enforced by the Harness platform.

## Full-Build Pipeline

The Full-Build Pipeline for a Go application must include the following stages, executed in this order:

### 1. Dependencies

Download and verify all module dependencies.

```bash
go mod download
go mod verify
```

### 2. Lint

Check code formatting, static analysis, and style conformance.

```bash
test -z "$(gofmt -l .)"
go vet ./...
golangci-lint run ./...
```

### 3. Test

Run all unit tests with the race detector enabled and generate a coverage report.

```bash
go test -race -covermode=atomic -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

### 4. Build

Compile a statically linked, reproducible binary.

```bash
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bin/app ./cmd/app
```

### 5. Security Scan

Scan the source code and its dependencies for known vulnerabilities.

```bash
govulncheck ./...
gosec ./...
```
