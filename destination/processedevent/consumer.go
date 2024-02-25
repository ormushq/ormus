package processedevent

import "github.com/ormushq/ormus/event"

type Consumer interface {
	Consume() (<-chan event.ProcessedEvent, error)
	Close() error
}
