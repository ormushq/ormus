package main

import (
	"github.com/google/uuid"
	adapter "github.com/ormushq/ormus/adapter/rabbitmq"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/ormushq/ormus/source/adapter/rabbitmq"
	"log"
)

func main() {
	cfg := config.C()
	r, err := adapter.New(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	pub := rabbitmq.NewRabbitMQAdapter(r)

	for range 200 {
		msg := source.NewSourceEvent{
			ProjectId: uuid.New().String(),
			OwnerId:   uuid.New().String(),
			WriteKey:  uuid.New().String(),
		}
		err := pub.Publish(cfg.Source.NewSourceEventName, []byte(encoder.EncodeNewSourceEvent(msg)))
		if err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}

	}

}
