package logger

import (
	"io"
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultFilePath        = "logs/json.log"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 30
)

type Config struct {
	FilePath         string
	UseLocalTime     bool
	FileMaxSizeInMB  int
	FileMaxAgeInDays int
}

var (
	L    *slog.Logger
	once = sync.Once{}
)

func init() {

	fileWriter := addSync(&lumberjack.Logger{
		Filename:  defaultFilePath,
		LocalTime: defaultUseLocalTime,
		MaxSize:   defaultFileMaxSizeInMB,
		MaxAge:    defaultFileAgeInDays,
	})
	L = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}),
	)
}

func New(cfg Config, opt *slog.HandlerOptions) {
	once.Do(func() {
		fileWriter := addSync(&lumberjack.Logger{
			Filename:  cfg.FilePath,
			LocalTime: cfg.UseLocalTime,
			MaxSize:   cfg.FileMaxSizeInMB,
			MaxAge:    cfg.FileMaxAgeInDays,
		})

		L = slog.New(
			slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), opt),
		)
	})
}
