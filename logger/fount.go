package logger

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/samber/lo"
)

var _ slog.Handler = (*fanoutHandler)(nil)

type fanoutHandler struct {
	handlers []slog.Handler
}

func fanout(handlers ...slog.Handler) slog.Handler {
	return &fanoutHandler{
		handlers: handlers,
	}
}

func (h *fanoutHandler) Enabled(ctx context.Context, l slog.Level) bool {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, l) {
			return true
		}
	}

	return false
}

func (h *fanoutHandler) Handle(ctx context.Context, r slog.Record) error {
	// @TODO: return multiple errors ?
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, r.Level) {
			err := try(func() error {
				return h.handlers[i].Handle(ctx, r.Clone())
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (h *fanoutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handers := lo.Map(h.handlers, func(h slog.Handler, _ int) slog.Handler {
		return h.WithAttrs(attrs)
	})

	return fanout(handers...)
}

func (h *fanoutHandler) WithGroup(name string) slog.Handler {
	handers := lo.Map(h.handlers, func(h slog.Handler, _ int) slog.Handler {
		return h.WithGroup(name)
	})

	return fanout(handers...)
}

func try(callback func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("unexpected error: %+v", r)
			}
		}
	}()

	err = callback()

	return nil
}
