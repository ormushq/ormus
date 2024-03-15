package inmemorytaskmanager

import (
	"log"

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
	log.Printf("Task [%s] is published to In-Memory queue.", pe.ID())
	q.events <- pe

	return nil
}
