package logger

import (
	"context"
	"log/slog"
	"os"
)

var L = slog.New(newLoggerHandler(nil))

type handler struct {
	h slog.Handler
}

func newLoggerHandler(opts *slog.HandlerOptions) *handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	return &handler{
		h: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     opts.Level,
			AddSource: true,
		}),
	}
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{h: h.h.WithAttrs(attrs)}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{h: h.h.WithGroup(name)}
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	return h.h.Handle(ctx, r)
}
