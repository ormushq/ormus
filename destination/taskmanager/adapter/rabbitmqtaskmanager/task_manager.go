package rabbitmqtaskmanager

import (
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/event"
)

type TaskPublisher struct {
	queue *Queue
}

func NewTaskPublisher(c dconfig.RabbitMQTaskManagerConnection, queueName string) *TaskPublisher {
	q := newQueue(c, queueName)

	return &TaskPublisher{
		queue: q,
	}
}

func (tm *TaskPublisher) Publish(pe event.ProcessedEvent) error {
	err := tm.queue.Enqueue(pe)
	if err != nil {
		return err
	}

	return nil
}
