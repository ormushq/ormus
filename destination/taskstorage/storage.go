package taskstorage

import (
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/event"
)

type Storage interface {
	CreateTask(pe event.ProcessedEvent) *entity.Task
	ChangeTaskStatus(TaskID string, newStatus string) error
	GetTaskByID(TaskID string) (error, *entity.Task)
}
