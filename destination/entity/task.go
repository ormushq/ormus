package entity

import (
	"github.com/ormushq/ormus/event"
	"time"
)

// Task represents a delivering processed event to third party integrations.
type Task struct {
	ID             string
	Name           string
	ProcessedEvent event.ProcessedEvent
	Timestamp      time.Time
}
