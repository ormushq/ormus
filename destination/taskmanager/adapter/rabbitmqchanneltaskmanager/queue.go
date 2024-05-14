package rabbitmqchanneltaskmanager

import (
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/event"
	"log/slog"
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
	fmt.Println(string(jsonEvent))
	q.channel <- jsonEvent

	return nil
}
