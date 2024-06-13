package slog

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/logger/newlogger/loggerparam"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
	"sync"
)

type Slog struct {
	config Config
	logger *slog.Logger
}

const fatalLevel = slog.Level(12)

var (
	DriverName       = "slog"
	once             = sync.Once{}
	slogLevelMapping = map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
		"fatal": fatalLevel,
	}
)

type Config struct {
	Level      string
	Filename   string
	Filepath   string
	LocalTime  bool
	MaxBackups int
	MaxSize    int
	MaxAge     int
}

func NewSlog(cfg Config) *Slog {
	logger := &Slog{config: cfg}
	logger.Init()

	return logger
}

func (s *Slog) getLogLevel() slog.Level {
	level, exists := slogLevelMapping[s.config.Level]
	if !exists {
		level = slog.LevelDebug
	}

	return level
}

func (s *Slog) Init() {
	once.Do(func() {
		fileWriter := &lumberjack.Logger{
			Filename:   s.config.Filename,
			LocalTime:  s.config.LocalTime,
			MaxSize:    s.config.MaxSize,
			MaxBackups: s.config.MaxBackups,
			MaxAge:     s.config.MaxAge,
		}
		logger := slog.New(
			slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{
				Level:       s.getLogLevel(),
				ReplaceAttr: ReplaceAttr,
			}),
		)
		s.logger = logger
	})
}

func ReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Value.String() == "ERROR+4" {
		a.Value = slog.StringValue("FATAL")
	}

	return a
}

func (s *Slog) Debug(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	s.logger.LogAttrs(context.Background(), slog.LevelDebug, msg, params...)
}

func (s *Slog) Debugf(template string, args ...interface{}) {
	s.logger.Debug(fmt.Sprintf(template, args...))
}

func (s *Slog) Info(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	s.logger.LogAttrs(context.Background(), slog.LevelInfo, msg, params...)
}

func (s *Slog) Infof(template string, args ...interface{}) {
	s.logger.Info(fmt.Sprintf(template, args...))
}

func (s *Slog) Warn(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	s.logger.LogAttrs(context.Background(), slog.LevelWarn, msg, params...)
}

func (s *Slog) Warnf(template string, args ...interface{}) {
	s.logger.Warn(fmt.Sprintf(template, args...))
}

func (s *Slog) Error(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	s.logger.LogAttrs(context.Background(), slog.LevelError, msg, params...)
}

func (s *Slog) Errorf(template string, args ...interface{}) {
	s.logger.Error(fmt.Sprintf(template, args...))
}

func (s *Slog) Fatal(cat loggerparam.Category, sub loggerparam.SubCategory, msg string, extra map[loggerparam.ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	s.logger.LogAttrs(context.Background(), fatalLevel, msg, params...)
	os.Exit(1)
}

func (s *Slog) Fatalf(template string, args ...interface{}) {
	s.logger.Log(context.Background(), fatalLevel, fmt.Sprintf(template, args...))
	os.Exit(1)
}

func prepareLogInfo(cat loggerparam.Category, sub loggerparam.SubCategory, extra map[loggerparam.ExtraKey]interface{}) []slog.Attr {
	if extra == nil {
		extra = make(map[loggerparam.ExtraKey]interface{})
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub

	return logParamsToSlogParams(extra)
}
