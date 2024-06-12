package newlogger

import (
	"github.com/ormushq/ormus/logger/newlogger/adapter/slog"
	"github.com/ormushq/ormus/logger/newlogger/loggerparam"
	"sync"
)

type Config struct {
	Driver     string `koanf:"driver"`
	Level      string `koanf:"level"`
	Filepath   string `koanf:"filepath"`
	LocalTime  bool   `koanf:"local_time"`
	MaxBackups int    `koanf:"max_backups"`
	MaxSize    int    `koanf:"max_size"`
	MaxAge     int    `koanf:"max_ager"`
}

type Logger interface {
	Init()

	Debug(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

var (
	once   = sync.Once{}
	logger Logger
)

func NewLogger(cfg Config) Logger {
	once.Do(func() {
		if cfg.Driver == slog.DriverName {
			logger = slog.NewSlog(slog.Config{
				Level:      cfg.Level,
				Filename:   cfg.Filepath,
				Filepath:   cfg.Filepath,
				LocalTime:  cfg.LocalTime,
				MaxBackups: cfg.MaxBackups,
				MaxSize:    cfg.MaxSize,
				MaxAge:     cfg.MaxAge,
			})

			return
		}
		panic("logger not supported")
	})

	return logger
}

func L() Logger {
	if logger == nil {
		panic("you need to init logger first")
	}

	return logger
}
