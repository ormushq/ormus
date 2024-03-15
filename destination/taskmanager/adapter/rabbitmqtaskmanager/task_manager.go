package rabbitmqtaskmanager

import (
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/event"
)

type TaskManager struct {
	queue *Queue
}

func NewTaskManager(c dconfig.RabbitMQTaskManagerConnection, queueName string) *TaskManager {
	q := newQueue(c, queueName)

	return &TaskManager{
		queue: q,
	}
}

func (tm *TaskManager) Publish(pe event.ProcessedEvent) error {
	err := tm.queue.Enqueue(pe)
	if err != nil {
		return err
	}

	return nil
}
