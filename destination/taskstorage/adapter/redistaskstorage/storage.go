package redistaskstorage

import (
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/event"
)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s Storage) CreateTask(pe event.ProcessedEvent) *entity.Task {
	//TODO implement CreateTask Redis task storage
	return nil
}

func (s Storage) ChangeTaskStatus(TaskID string, newStatus string) error {
	//TODO implement ChangeTaskStatus Redis task storage
	return nil
}

func (s Storage) GetTaskByID(TaskID string) (error, *entity.Task) {
	//TODO implement GetTaskByID Redis task storage
	return nil, nil
}
