package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

const timeoutSeconds = 5

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"processed-events-exchange", // name
		"topic",                     // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	fakeIntegration := entity.Integration{
		ID:       "5",
		SourceID: "1",
		Metadata: entity.DestinationMetadata{
			ID:   "1",
			Name: "webhook",
			Slug: "webhook",
		},
	}
	// generate fake processedEvent
	pageName := "Home"
	pe := event.ProcessedEvent{
		SourceID:          "1",
		Integration:       fakeIntegration,
		MessageID:         "1",
		EventType:         "page",
		Name:              &pageName,
		Version:           1,
		SentAt:            time.Now(),
		ReceivedAt:        time.Now(),
		OriginalTimestamp: time.Now(),
		Timestamp:         time.Now(),
	}

	jpe, err := json.Marshal(pe)
	if err != nil {
		log.Panicf("Error: %e", err)
	}

	err = ch.PublishWithContext(ctx,
		"processed-events-exchange", // exchange
		"pe.webhook",                // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jpe,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("Publish new processed event.")
}
