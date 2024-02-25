package inmemorytaskmanager

import (
	"fmt"
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/destination/taskstorage"
)

type TaskManager struct {
	Queue       *Queue
	taskStorage taskstorage.Storage
}

func NewTaskManager(ts taskstorage.Storage) *TaskManager {
	q := NewQueue()
	return &TaskManager{
		Queue:       q,
		taskStorage: ts,
	}
}

func (tm *TaskManager) SendToQueue(t *entity.Task) error {

	err := tm.Queue.Enqueue(t)
	if err != nil {
		fmt.Println("Enqueue Error : ", err)
		return err
	}

	return nil
}
