package logger

import (
	"github.com/ormushq/ormus/pkg/trace"
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// TODO: should we transfer these constants to the default config struct also? or not?
const (
	defaultFilePath        = "logs/logs.json"
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

var l *slog.Logger

func init() {
	fileWriter := &lumberjack.Logger{
		Filename:  defaultFilePath,
		LocalTime: defaultUseLocalTime,
		MaxSize:   defaultFileMaxSizeInMB,
		MaxAge:    defaultFileAgeInDays,
	}
	l = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}),
	)
}

func L() *slog.Logger {
	return l
}
func WithGroup(groupName string) *slog.Logger {
	t := trace.Parse()
	return l.With(slog.Group(groupName,
		slog.String("path", t.File),
		slog.Int("line", t.Line),
		slog.String("function", t.Function),
	))
}

func New(cfg Config, opt *slog.HandlerOptions) *slog.Logger {
	fileWriter := &lumberjack.Logger{
		Filename:  cfg.FilePath,
		LocalTime: cfg.UseLocalTime,
		MaxSize:   cfg.FileMaxSizeInMB,
		MaxAge:    cfg.FileMaxAgeInDays,
	}

	logger := slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), opt),
	)

	return logger
}
