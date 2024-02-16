package queue

import (
	"github.com/ormushq/ormus/destination/entity"
	"time"
)

// Job represents a task to be processed by a worker.
type Job struct {
	ID             int
	Topic          string
	ProcessedEvent entity.ProcessedEvent
	Timestamp      time.Time
}

// JobQueue defines the interface for a job queue.
type JobQueue interface {
	Enqueue(job Job)
	Dequeue(topic string) Job
}
