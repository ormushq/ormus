package logger

import (
	"io"
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
	L    *slog.Logger
	once = sync.Once{}
)

func init() {
	defaultConfig := Config{
		FinePath:         "logs/logs.json",
		UseLocalTime:     false,
		FileMaxSizeInMB:  10,
		FineMaxAgeInDays: 30,
	}

	fileWriter := addSync(&lumberjack.Logger{
		Filename:  defaultConfig.FinePath,
		LocalTime: defaultConfig.UseLocalTime,
		MaxSize:   defaultConfig.FileMaxSizeInMB,
		MaxAge:    defaultConfig.FineMaxAgeInDays,
	})
	L = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}),
	)
}

func New(cfg Config, opt *slog.HandlerOptions) {
	once.Do(func() {
		fileWriter := addSync(&lumberjack.Logger{
			Filename:  cfg.FinePath,
			LocalTime: cfg.UseLocalTime,
			MaxSize:   cfg.FileMaxSizeInMB,
			MaxAge:    cfg.FineMaxAgeInDays,
		})

		L = slog.New(
			slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), opt),
		)
	})
}
