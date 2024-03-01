package main

import (
	"context"
	"encoding/json"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

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
		"processed_events_exchange", // name
		"topic",                     // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//generate fake processedEvent
	pageName := "Home"
	//pe := event.ProcessedEvent{
	//	SourceID:          "1",
	//	Integration:       entity.Integration{ID: "2"},
	//	MessageID:         "2",
	//	EventType:         "page",
	//	Name:              &pageName,
	//	Version:           1,
	//	SentAt:            time.Now(),
	//	ReceivedAt:        time.Now(),
	//	OriginalTimestamp: time.Now(),
	//	Timestamp:         time.Now(),
	//}
	pe := event.ProcessedEvent{
		SourceID: "2",
		Integration: entity.Integration{
			Config: entity.WebhookConfig{
				Headers: []entity.Header{
					{Key: "Authorization", Value: "Basic MY_BASIC_AUTH_TOKEN"},
					{Key: "Content-Type", Value: "MY_CONTENT_TYPE"},
				},
				Payload: []entity.Payload{
					{Key: "test1", Value: "value test1"},
					{Key: "test2", Value: "value test2"},
					{Key: "test3", Value: "value test3"},
				},
				Method: entity.GETWebhookMethod,
				Url:    "https://eoc0z7vqfxu6io.m.pipedream.net",
			},
		},
		MessageID:         "12",
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
		log.Panicf("Error:", err)
	}

	err = ch.PublishWithContext(ctx,
		"processed_events_exchange", // exchange
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
