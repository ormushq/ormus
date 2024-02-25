package faketaskstorage

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
	//TODO implement me
	return nil
}

func (s Storage) ChangeTaskStatus(TaskID string, newStatus string) error {
	//TODO implement me
	return nil
}

func (s Storage) GetTaskByID(TaskID string) (error, *entity.Task) {
	//TODO implement me
	return nil, nil
}
