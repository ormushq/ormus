package logger

import (
	"errors"
	"io"
	"log/slog"
	"os"

	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/pkg/trace"
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

func LogError(err error) {
	if err == nil {
		return
	}
	var rErr richerror.RichError
	if errors.As(err, &rErr) {
		L().Error(rErr.Message())
	}
}

func WithGroup(groupName string) *slog.Logger {
	t := trace.Parse()

	return l.With(slog.String("group", groupName)).With(slog.Group("trace",
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

func SetLevel(modeLevel slog.Level) {
	fileWriter := &lumberjack.Logger{
		Filename:  defaultFilePath,
		LocalTime: defaultUseLocalTime,
		MaxSize:   defaultFileMaxSizeInMB,
		MaxAge:    defaultFileAgeInDays,
	}
	l = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{
			Level: modeLevel,
		}),
	)
}
