# GoPlate Agent Context

This guide provides essential dev context for automated agents working in this repository (GoPlate — a Go-based REST API boilerplate).

## 1. Build, Lint, and Test
- **Build**: `make build` or `go build -o server main.go`
- **Run**: `make run` (builds then runs)
- **Dev (hot reload)**: `make dev` (needs reflex)
- **Format**: `make fmt` or `go fmt ./...`
- **Test all**: `make test` or `go test ./...`
- **Test (single file)**: `go test ./path/to/pkg/file_test.go`
- **Test Coverage**: `make test-coverage`
- **Linting**: Uses `go fmt` (strict gofmt style); no golangci-lint config by default

## 2. Code Style Guidelines
- **Imports**: Standard library, then third-party, then local imports (grouped, no blank lines between groups)
- **Formatting**: Always use `gofmt`; tabs not spaces
- **Types**: Prefer explicit struct field types; use `interface{}` or `any` sparingly (replace where modern Go suggests)
- **Naming**: `CamelCase` for structs/interfaces, `camelCase` for vars, `ALL_CAPS` for constants; exported symbols are Capitalized
- **Error handling**: Return Go `error` type; wrap or enrich errors with context (e.g., `fmt.Errorf`); log errors with `logs.Error`/`logs.Fatal`
- **JSON struct tags**: Always used for API inputs/outputs
- **Tests**: Standard `_test.go` files; assertions via Go testing or testify if added
- **Env/config**: Managed via `env/` package and `.env` file

## 3. Project Structure
- App code is under `cmd/`, `pkg/`, `domains/`, `router/`
- DB/migrations/seed in `db/`, utility code in `pkg/utils/`

No Cursor or Copilot rules present. If new lint, type, or style guidelines are added, update this file accordingly.