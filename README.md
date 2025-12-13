# go-exp

Small, experimental helpers for Go that lean on generics to reduce boilerplate
without piling on dependencies. Expect the surface area to stay small and to
evolve as patterns prove useful across projects.

## Installation

```
go get github.com/min0625/go-exp
```

## Usage

APIs live in the root package and future subpackages. Browse the documentation
with `go doc github.com/min0625/go-exp` or `go doc` on specific identifiers once
they land.

## Principles

- Keep helpers narrowly scoped and composable.
- Prefer zero dependencies and standard library first.
- Use type parameters where they improve clarity and safety.

## Development

- Requires Go 1.25 or newer (see go.mod).
- Run tests with `go test ./...`.
