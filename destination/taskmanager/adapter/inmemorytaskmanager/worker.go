package inmemorytaskmanager

import (
	"sync"

	"github.com/ormushq/ormus/destination/worker"
)

type Worker struct {
	TaskManager *TaskManager
	Handler     worker.TaskHandler
}

func NewWorker(tm *TaskManager, h worker.TaskHandler) *Worker {
	return &Worker{
		TaskManager: tm,
		Handler:     h,
	}
}

func (w *Worker) Run(_ <-chan bool, _ *sync.WaitGroup) error {
	// todo implement in-memory worker

	return nil
}
