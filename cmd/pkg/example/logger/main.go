package main

import (
	"github.com/ormushq/ormus/logger/newlogger"
	"github.com/ormushq/ormus/logger/newlogger/loggerparam"
	"time"
)

func main() {
	MaxBackups := 10
	MaxSize := 10
	MaxAge := 30

	newlogger.NewLogger(newlogger.Config{
		Driver:     "slog",
		Level:      "debug",
		Filepath:   "logs/logs.json",
		LocalTime:  false,
		MaxBackups: MaxBackups,
		MaxSize:    MaxSize,
		MaxAge:     MaxAge,
	})

	extraKeyValue := "extra key not defined"

	newlogger.NewLog("Test Debug").
		WithCategory(loggerparam.CategoryNotDefined).
		WithSubCategory(loggerparam.SubCategoryNotDefined).
		With(loggerparam.ExtraKeyNotDefined, extraKeyValue).
		WithTrace().
		Debug()
	newlogger.L().Debugf("Test Debugf %v", time.Now())

	newlogger.NewLog("Test Info").
		WithCategory(loggerparam.CategoryNotDefined).
		WithSubCategory(loggerparam.SubCategoryNotDefined).
		With(loggerparam.ExtraKeyNotDefined, extraKeyValue).
		WithTrace().
		Info()
	newlogger.L().Infof("Test Infof %v", time.Now())

	newlogger.NewLog("Test Warn").
		WithCategory(loggerparam.CategoryNotDefined).
		WithSubCategory(loggerparam.SubCategoryNotDefined).
		With(loggerparam.ExtraKeyNotDefined, extraKeyValue).
		WithTrace().
		Warn()
	newlogger.L().Warnf("Test Warnf %v", time.Now())

	newlogger.NewLog("Test Error").
		WithCategory(loggerparam.CategoryNotDefined).
		WithSubCategory(loggerparam.SubCategoryNotDefined).
		With(loggerparam.ExtraKeyNotDefined, extraKeyValue).
		WithTrace().
		Error()
	newlogger.L().Errorf("Test Errorf %v", time.Now())

	newlogger.NewLog("Test Fatal").
		WithCategory(loggerparam.CategoryNotDefined).
		WithSubCategory(loggerparam.SubCategoryNotDefined).
		With(loggerparam.ExtraKeyNotDefined, extraKeyValue).
		WithTrace().
		Fatal()
	newlogger.L().Fatalf("Test Fatalf %v", time.Now())
}
