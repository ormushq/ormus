package inmemorytaskmanager

import (
	"github.com/ormushq/ormus/event"
)

// Queue represents an in-memory job queue.
type Queue struct {
	events chan event.ProcessedEvent
}

func NewQueue() *Queue {
	channelSize := 100
	events := make(chan event.ProcessedEvent, channelSize)

	return &Queue{
		events: events,
	}
}

// Enqueue adds a new job to the in-memory queue.
func (q *Queue) Enqueue(pe event.ProcessedEvent) error {
	q.events <- pe

	return nil
}
