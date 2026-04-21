# AGENTS.md

## Project Overview

go-exp is a small, experimental collection of generic helpers for Go.

- Root package: `goexp` (documentation-only package declaration in `doc.go`)
- Subpackages:
  - `errorx`: error helpers
  - `iterx`: iterator helpers (`iter.Seq`)
  - `ptr`: pointer helpers
  - `slicex`: slice helpers, composes `iterx`
  - `slogx`: `log/slog` helpers
  - `util`: miscellaneous utilities

Design principles from the repository:

- Keep helpers small and composable
- Prefer standard library and zero external runtime dependencies
- Use generics where they improve clarity and safety

## Environment and Setup

- Required Go version: `1.25.7` (from `go.mod`)
- Module path: `github.com/min0625/go-exp`

Initial setup:

```bash
go version
go test ./...
```

## Development Workflow

Use Makefile targets as the canonical workflow:

```bash
make lint  # run golangci-lint checks
make test  # go test -v -race -failfast ./...
make check # lint + test
```

Notes:

- `make lint` and `make test` are validated and runnable in this repository.
- `make fix` applies auto-fixes from golangci-lint where available.

## Testing Instructions

Run the full suite:

```bash
make test
```

Direct equivalent:

```bash
go test -v -race -failfast ./...
```

Current test layout:

- Example tests exist in `slicex/example_iter_test.go`.
- Many packages currently have no test files; keep this expected state in mind.

When changing behavior, add or update tests in the affected package.

## Code Style and Linting

Lint/format configuration is in `.golangci.yaml`.

- The repo uses golangci-lint v2-style config (`version: "2"`).
- Formatters include: `gofmt`, `gofumpt`, `goimports`, `gci`, `golines`.
- Enabled linters include `govet`, `staticcheck`, `gosec`, `errcheck`, and others.

Important guardrails:

- Avoid non-standard replacements blocked by `depguard` (for example `github.com/pkg/errors`, `golang.org/x/exp/slices`).
- Follow stdlib-first approach documented in README and lint policy.
- Use `any` instead of `interface{}` (enforced in formatter rewrite rules).

## CI and Pull Requests

GitHub Actions workflows:

- `.github/workflows/go.yml`: lint + test on pull requests
- `.github/workflows/codeql.yml`: CodeQL analysis for Go

Before opening or updating a PR, run:

```bash
make check
```

Minimum expectation for agent-generated changes:

1. Keep changes scoped to the target package(s).
2. Run lint and tests locally.
3. Update/add tests when behavior changes.

## Security Notes

- This project runs `gosec` via golangci-lint and CodeQL in CI.
- Do not introduce secrets, tokens, or credentials in source/tests.
- For randomness, use `math/rand/v2` only for non-cryptographic contexts unless requirements explicitly call for cryptographic randomness.

## Agent-Specific Guidance

When modifying this repository:

1. Prefer the smallest API-compatible change.
2. Preserve package boundaries and naming conventions (`*x` helper packages).
3. Keep helper functions generic and dependency-light.
4. If adding new helper families, place them in a focused subpackage rather than expanding unrelated files.
5. Update this file when commands, CI, or tooling expectations change.
