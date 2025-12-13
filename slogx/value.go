// Package slogx provides extensions for the slog package.
package slogx

import (
	"log/slog"
)

// FuncValuer wraps a function that returns slog.Value into slog.LogValuer.
func FuncValuer(f func() slog.Value) slog.LogValuer {
	return funcValuer(f)
}

// ValueFunc is a function type that returns slog.Value.
type ValueFunc = func() slog.Value

type funcValuer func() slog.Value

var _ slog.LogValuer = funcValuer(nil)

func (f funcValuer) LogValue() slog.Value {
	return f()
}
