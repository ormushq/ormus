package broker

import "github.com/ormushq/ormus/event"

type Consumer interface {
	Consume(topic string) <-chan event.CoreEvent
}
