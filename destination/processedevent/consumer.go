package processedevent

import "github.com/ormushq/ormus/event"

type Consumer interface {
	Consume(done <-chan bool) (<-chan event.ProcessedEvent, error)
	Close() error
}
