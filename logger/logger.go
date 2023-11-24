package logger

import (
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	fileName         = "logs/logs.json"
	fileMaxSizeInMB  = 10
	fileMaxAgeInDays = 30
)

var L *slog.Logger

var once = sync.Once{}

func init() {
	once.Do(func() {
		fileWriter := addSync(&lumberjack.Logger{
			Filename:  fileName,
			LocalTime: false,
			MaxSize:   fileMaxSizeInMB,
			MaxAge:    fileMaxAgeInDays,
		})

		L = slog.New(
			fanout(
				slog.NewJSONHandler(fileWriter, &slog.HandlerOptions{}), // pass to first handler: logstash over tcp
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),  // then to second handler: stderr
			),
		)
	})
}
