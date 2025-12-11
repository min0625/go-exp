package slogx

import (
	"context"
	"log/slog"
)

// ctxAttrsKey is the private key type used for storing slog attributes in context.
type ctxAttrsKey struct{}

// WithAttrs appends slog attributes into context for later automatic logging.
func WithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	if len(attrs) == 0 {
		return ctx
	}

	existing, _ := ctx.Value(ctxAttrsKey{}).([]slog.Attr)

	// Copy to avoid modifying underlying slice if shared.
	newSlice := make([]slog.Attr, 0, len(existing)+len(attrs))
	newSlice = append(newSlice, existing...)
	newSlice = append(newSlice, attrs...)

	return context.WithValue(ctx, ctxAttrsKey{}, newSlice)
}

// NewContextHandler creates a handler that enriches records with ctx fields.
func NewContextHandler(inner slog.Handler) *ContextHandler {
	return &ContextHandler{inner: inner}
}

// ContextHandler wraps an underlying slog.Handler and injects context fields.
type ContextHandler struct {
	inner slog.Handler
}

var _ slog.Handler = (*ContextHandler)(nil)

// Enabled proxies to the inner handler.
func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.inner.Enabled(ctx, level)
}

// Handle enriches the record with context-provided attributes before delegating.
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	attrs, _ := ctx.Value(ctxAttrsKey{}).([]slog.Attr)
	if len(attrs) == 0 {
		return h.inner.Handle(ctx, r)
	}
	// Create a new record to avoid mutating the original.
	newRec := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	for _, a := range attrs {
		newRec.AddAttrs(a)
	}
	// Preserve original attributes.
	r.Attrs(func(a slog.Attr) bool {
		newRec.AddAttrs(a)
		return true
	})

	return h.inner.Handle(ctx, newRec)
}

// WithAttrs returns a new ContextHandler with attributes applied to the inner handler.
func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{inner: h.inner.WithAttrs(attrs)}
}

// WithGroup returns a new ContextHandler with group applied on the inner handler.
func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{inner: h.inner.WithGroup(name)}
}
