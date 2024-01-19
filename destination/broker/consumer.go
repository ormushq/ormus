package broker

import (
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
)

type ConsumerConfig struct {
	Topic string `koanf:"consumer_topic"`
}

type ProcessedEvent struct {
	Event       event.CoreEvent
	Integration []entity.Integration
}

// TODO: implement a hybrid solution (get topic from config file and constructor method).
type Consumer interface {
	Consume() <-chan ProcessedEvent
	Close() error
}
