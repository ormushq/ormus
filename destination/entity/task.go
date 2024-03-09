package entity

import (
	"time"

	"github.com/ormushq/ormus/event"
)

type TaskStatus uint8

const (
	NotExists                     TaskStatus = 1
	FailedInIntegrationHandler    TaskStatus = 2
	SuccessInIntegrationHandler   TaskStatus = 3
	SuccessInUpdateDeliveryStatus TaskStatus = 4
	FailedInUpdateDeliveryStatus  TaskStatus = 5
)

// Task represents a delivering processed event to third party integrations.
type Task struct {
	ID             string
	Name           string
	Status         *TaskStatus
	ProcessedEvent event.ProcessedEvent
	Timestamp      time.Time
}
