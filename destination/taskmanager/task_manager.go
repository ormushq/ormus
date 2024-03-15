package taskmanager

import (
	"sync"

	"github.com/ormushq/ormus/event"
)

type Publisher interface {
	Publish(event event.ProcessedEvent) error
}

type Consumer interface {
	Consume(done <-chan bool, wg *sync.WaitGroup) (<-chan event.ProcessedEvent, error)
}
