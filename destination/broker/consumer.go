package broker

import "github.com/ormushq/ormus/event"

type ConsumerConfig struct {
	Topic string `koanf:"consumer_topic"`
}

// TODO: implement a hybrid solution (get topic from config file and constructor method).
type Consumer interface {
	Consume() <-chan event.CoreEvent
	Close() error
}
