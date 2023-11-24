package logger

import (
	"io"
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Filename    string `koanf:"filename"`
	MaxSize     int    `koanf:"maxsize"`
	MaxBackups  int    `koanf:"maxbackups"`
	MaxAge      int    `koanf:"maxage"`
	Compress    bool   `koanf:"compress"`
	LocalTime   bool   `koanf:"localtime"`
	LogLevel    string `koanf:"log_level"`
	AddSource   bool   `koanf:"add_source_log"`
	HandlerType string `koanf:"handler_type"`
}

var Logger *slog.Logger

var once = sync.Once{}

func InitLogger(cfg Config) {

	once.Do(func() {

		var logLevel slog.Level

		switch cfg.LogLevel {

		case "DEBUG":
			logLevel = slog.LevelDebug

		case "WARNING":
			logLevel = slog.LevelWarn

		case "ERROR":
			logLevel = slog.LevelError

		default:
			// INFO level
			logLevel = slog.LevelInfo
		}

		fileWriter := &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  cfg.LocalTime,
		}

		multiWriter := io.MultiWriter(os.Stdout, fileWriter)

		switch cfg.HandlerType {

		case "json":

			Logger = slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: logLevel, AddSource: cfg.AddSource}))

		default:
			// text handler
			Logger = slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: logLevel, AddSource: cfg.AddSource}))

		}

		slog.SetDefault(Logger)

	})
}
