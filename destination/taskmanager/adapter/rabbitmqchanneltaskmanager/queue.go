package rabbitmqchanneltaskmanager

import (
	"encoding/json"
	"log/slog"

	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
)

type Queue struct {
	channel chan<- []byte
}

func newQueue(inputChannel chan<- []byte) *Queue {
	return &Queue{
		channel: inputChannel,
	}
}

func (q *Queue) Enqueue(pe event.ProcessedEvent) error {
	// Convert Processed event to json
	jsonEvent, err := json.Marshal(pe)
	if err != nil {
		slog.Error("Error:", err)

		return err
	}
	logger.L().Debug(string(jsonEvent))
	q.channel <- jsonEvent

	return nil
}
