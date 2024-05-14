package rabbitmqchanneltaskmanager

import (
	"github.com/ormushq/ormus/event"
)

type TaskPublisher struct {
	queue *Queue
}

func NewTaskPublisher(inputChannel chan<- []byte) *TaskPublisher {
	q := newQueue(inputChannel)

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
