package main

import (
	"github.com/google/uuid"
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/ormushq/ormus/source/adapter/rabbitmq"
)

func main() {

	pub, _ := rabbitmq.NewRabbitMQAdapter("amqp://guest:guest@localhost:5672/")

	for range 200 {
		msg := source.NewSourceEvent{
			ProjectId: uuid.New().String(),
			OwnerId:   uuid.New().String(),
			WriteKey:  uuid.New().String(),
		}
		err := pub.Publish("new-source-event", []byte(encoder.EncodeNewSourceEvent(msg)))
		if err != nil {
			panic(err)
		}

	}

}
