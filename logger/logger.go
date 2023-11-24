package logger

import (
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	FinePath         string
	UseLocalTime     bool
	FileMaxSizeInMB  int
	FineMaxAgeInDays int
}

var (
	l    *slog.Logger
	once = sync.Once{}
)

func Logger(cfg Config, opt *slog.HandlerOptions) *slog.Logger {
	once.Do(func() {
		fileWriter := addSync(&lumberjack.Logger{
			Filename:  cfg.FinePath,
			LocalTime: cfg.UseLocalTime,
			MaxSize:   cfg.FileMaxSizeInMB,
			MaxAge:    cfg.FineMaxAgeInDays,
		})

		l = slog.New(
			fanout(
				slog.NewJSONHandler(fileWriter, opt),
				slog.NewJSONHandler(os.Stdout, opt),
			),
		)
	})

	return l
}
