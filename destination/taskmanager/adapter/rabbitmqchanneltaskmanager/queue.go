package rabbitmqchanneltaskmanager

import (
	"encoding/json"
	"github.com/ormushq/ormus/pkg/channel"
	"log/slog"

	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
)

type Queue struct {
	channel chan<- channel.Message
}

func newQueue(inputChannel chan<- channel.Message) *Queue {
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
	q.channel <- channel.Message{
		Body: jsonEvent,
	}

	return nil
}
