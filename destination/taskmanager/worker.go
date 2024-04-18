package taskmanager

import (
	"context"
	"sync"

	"github.com/ormushq/ormus/event"
)

// Worker represents a worker that processes jobs of queue.
type Worker interface {
	Run(done <-chan bool, wg *sync.WaitGroup)
}

type TaskHandler interface {
	HandleTask(ctx context.Context, newEvent event.ProcessedEvent) error
}
