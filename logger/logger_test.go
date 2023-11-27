package logger_test

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"testing"

	"github.com/ormushq/ormus/logger"
)

func TestLogger(t *testing.T) {
	cfg := logger.Config{
		FilePath:         "./logs.json",
		UseLocalTime:     false,
		FileMaxSizeInMB:  10,
		FileMaxAgeInDays: 1,
	}
	opt := slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				// remove time because it makes test wrong `time is so fast :)`
				return slog.Attr{}
			}

			return a
		},
	}
	logger.New(cfg, &opt)

	tests := []struct {
		f    func()
		want string
	}{
		{
			f: func() {
				logger.L.Info("INFO", "key", "value")
			},
			want: `{"level":"INFO","msg":"INFO","key":"value"}`,
		},
		{
			f: func() {
				logger.L.Warn("WARN", "key", "value")
			},
			want: `{"level":"WARN","msg":"WARN","key":"value"}`,
		},
		{
			f: func() {
				logger.L.Error("ERROR", "key", "value")
			},
			want: `{"level":"ERROR","msg":"ERROR","key":"value"}`,
		},
		{
			f: func() {
				logger.L.With(
					slog.Group("user",
						slog.String("id", "user-123"),
					),
				).
					With("environment", "dev").
					With("error", fmt.Errorf("an error")).
					Error("A message")
			},
			want: `{"level":"ERROR","msg":"A message","user":{"id":"user-123"},"environment":"dev","error":"an error"}`,
		},
	}

	// first run logs
	for _, test := range tests {
		test.f()
	}

	f, err := os.Open(cfg.FilePath)
	defer f.Close()
	defer os.Remove(cfg.FilePath)
	if err != nil {
		t.Fatalf("can't open file: %s", err)
	}

	scanner := bufio.NewScanner(f)

	var logs []string

	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if test.want != logs[i] {
				t.Fatalf("want: %+v, got: %+v", test.want, logs[i])
			}
		})
	}
}
