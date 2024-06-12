package slog

import (
	"github.com/ormushq/ormus/logger/newlogger/loggerparam"
	"log/slog"
)

func logParamsToSlogParams(keys map[loggerparam.ExtraKey]interface{}) []any {
	params := make([]any, 0, len(keys))

	for k, v := range keys {
		params = append(params, slog.Any(string(k), v))
	}

	return params
}
