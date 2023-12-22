package broker

import "github.com/ormushq/ormus/event"

type ConsumerConfig struct {
	Topic string `koanf:"consumer_topic"`
}

type Consumer interface {
	Consume() <-chan event.CoreEvent
	Close() // TODO: Do we need to worry about resource leak?
}
