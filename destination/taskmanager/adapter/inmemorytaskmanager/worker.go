package inmemorytaskmanager

import (
	"sync"

	"github.com/ormushq/ormus/destination/taskmanager"
)

type Worker struct {
	TaskManager *TaskManager
	Handler     taskmanager.TaskHandler
}

func NewWorker(tm *TaskManager, h taskmanager.TaskHandler) *Worker {
	return &Worker{
		TaskManager: tm,
		Handler:     h,
	}
}

func (w *Worker) Run(_ <-chan bool, _ *sync.WaitGroup) error {
	// todo implement in-memory worker

	return nil
}
