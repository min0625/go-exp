package slogx

import (
	"log/slog"
)

func FuncValuer(f func() slog.Value) slog.LogValuer {
	return funcValuer(f)
}

type ValueFunc = func() slog.Value

type funcValuer func() slog.Value

var _ slog.LogValuer = funcValuer(nil)

func (f funcValuer) LogValue() slog.Value {
	return f()
}
