package inmemorytaskmanager

import (
	"log/slog"

	"github.com/ormushq/ormus/event"
)

type TaskManager struct {
	Queue *Queue
}

func New() *TaskManager {
	q := NewQueue()

	return &TaskManager{
		Queue: q,
	}
}

func (tm *TaskManager) Publish(e event.ProcessedEvent) error {
	// send task to queue
	err := tm.Queue.Enqueue(e)
	if err != nil {
		slog.Error("enqueue Error : ", err)

		return err
	}

	return nil
}
