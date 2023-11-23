package logger

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"sync"
)

var Logger = slog.New(newHandler(nil))

const (
	reset      = "\033[0m"
	timeFormat = "[02/01/2006 15:04]"

	lightRed   = 91
	yellow     = 33
	cyan       = 36
	darkGray   = 90
	lightGreen = 92
	lightGray  = 37
	white      = 97
)

func colorize(colorCode int, level string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), level, reset)
}

type handler struct {
	h slog.Handler
	b *bytes.Buffer
	m *sync.Mutex
}

func newHandler(opts *slog.HandlerOptions) *handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	b := &bytes.Buffer{}
	return &handler{
		b: b,
		h: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   true,
			ReplaceAttr: overwriteDefaults(opts.ReplaceAttr),
		}),
		m: &sync.Mutex{},
	}
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{h: h.h.WithAttrs(attrs), b: h.b, m: h.m}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{h: h.h.WithGroup(name), b: h.b, m: h.m}
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = colorize(darkGray, level)
	case slog.LevelInfo:
		level = colorize(cyan, level)
	case slog.LevelWarn:
		level = colorize(yellow, level)
	case slog.LevelError:
		level = colorize(lightRed, level)
	}

	bytes, err := h.appendAttrs(ctx, r)
	if err != nil {
		return err
	}

	fmt.Println(
		colorize(lightGray, r.Time.Format(timeFormat)),
		level,
		colorize(lightGreen, r.Message),
		colorize(white, string(bytes)),
	)

	return nil
}

func (h *handler) appendAttrs(ctx context.Context, r slog.Record) ([]byte, error) {
	h.m.Lock()
	defer func() {
		h.b.Reset()
		h.m.Unlock()
	}()

	if err := h.h.Handle(ctx, r); err != nil {
		return nil, fmt.Errorf("error when calling inner handler's Handle: %w", err)
	}

	res := bytes.Trim(h.b.Bytes(), "\n")
	return res, nil
}

func overwriteDefaults(next func([]string, slog.Attr) slog.Attr) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey ||
			a.Key == slog.LevelKey ||
			a.Key == slog.MessageKey {
			return slog.Attr{}
		}
		if next == nil {
			return a
		}
		return next(groups, a)
	}
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
