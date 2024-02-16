package handler

import "github.com/ormushq/ormus/destination/queue"

// IntegrationHandler defines the interface for a topic handler.
type IntegrationHandler interface {
	Handle(job *queue.Job)
	GetTopic() string
}
