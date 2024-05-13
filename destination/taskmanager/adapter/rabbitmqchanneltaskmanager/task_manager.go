package rabbitmqchanneltaskmanager

import (
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/event"
	"sync"
)

type TaskPublisher struct {
	queue *Queue
}

func NewTaskPublisher(done <-chan bool, wg *sync.WaitGroup, c dconfig.RabbitMQTaskManagerConnection, queueName string, reconnectSecond int) *TaskPublisher {
	q := newQueue(done, wg, c, queueName, reconnectSecond)

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
