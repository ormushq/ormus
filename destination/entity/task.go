package entity

import (
	"github.com/ormushq/ormus/event"
	"time"
)

type TaskStatus uint8

const (
	NOT_EXISTS                        TaskStatus = 1
	FAILED_IN_INTEGRATION_HANDLER     TaskStatus = 2
	SUCCESS_IN_INTEGRATION_HANDLER    TaskStatus = 3
	SUCCESS_IN_UPDATE_DELIVERY_STATUS TaskStatus = 4
	FAILDE_IN_UPDATE_DELIVERY_STATUS  TaskStatus = 5
)

// Task represents a delivering processed event to third party integrations.
type Task struct {
	ID             string
	Name           string
	Status         *TaskStatus
	ProcessedEvent event.ProcessedEvent
	Timestamp      time.Time
}
