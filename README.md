# go-exp

[![Go Reference](https://pkg.go.dev/badge/github.com/min0625/go-exp.svg)](https://pkg.go.dev/github.com/min0625/go-exp)

Small, experimental helpers for Go that lean on generics to reduce boilerplate without piling on dependencies.

## Installation

```sh
go get github.com/min0625/go-exp
```

## Packages

| Package | Description |
|---------|-------------|
| [`errorx`](https://pkg.go.dev/github.com/min0625/go-exp/errorx) | Error handling helpers — generic `AsType`, `Unwrapper`, `MultiUnwrapper` |
| [`iterx`](https://pkg.go.dev/github.com/min0625/go-exp/iterx) | `iter.Seq` helpers — `Map`, `Filter` |
| [`ptr`](https://pkg.go.dev/github.com/min0625/go-exp/ptr) | Pointer helpers — `New`, `Value`, `ValueOr`, `Clone` |
| [`slicex`](https://pkg.go.dev/github.com/min0625/go-exp/slicex) | Slice helpers — `Map`, `Filter`, `FirstOr`, `Last`, `LastOr` |
| [`slogx`](https://pkg.go.dev/github.com/min0625/go-exp/slogx) | `log/slog` helpers — `FuncValuer` |
| [`util`](https://pkg.go.dev/github.com/min0625/go-exp/util) | Miscellaneous utilities — `PickOne` |

## Examples

```go
// ptr: take the address of a literal
p := ptr.New(42) // *int

// slicex: transform and filter slices
names := slicex.Map([]int{1, 2, 3}, strconv.Itoa) // ["1", "2", "3"]
evens := slicex.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }) // [2, 4]

// errorx: type-safe error unwrapping
if e, ok := errorx.AsType[*MyError](err); ok { ... }

// slogx: lazy log value evaluation
attr := slog.Any("key", slogx.FuncValuer(func() slog.Value { return slog.StringValue(compute()) }))
```

## Principles

- Keep helpers narrowly scoped and composable.
- Prefer zero dependencies and standard library first.
- Use type parameters where they improve clarity and safety.

## Development

Requires Go 1.25 or newer.

```sh
make test   # go test -v -race -failfast ./...
make lint   # run golangci-lint
make check  # lint + test
```
