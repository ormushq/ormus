package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	amqp "github.com/rabbitmq/amqp091-go"
)

const timeoutSeconds = 5

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	rmqConsumerConnConfig := config.C().Destination.RabbitMQConsumerConnection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rmqConsumerConnConfig.User,
		rmqConsumerConnConfig.Password, rmqConsumerConnConfig.Host, rmqConsumerConnConfig.Port))

	defer func(conn *amqp.Connection) {
		err = conn.Close()
		failOnError(err, "Failed to close RabbitMQ connection")
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err = ch.Close()
		failOnError(err, "Failed to close channel")
	}(ch)

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
		ID:       "10",
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
		SourceID: "2",
		Integration: entity.Integration{
			Config: webhookintegration.WebhookConfig{
				Headers: []webhookintegration.Header{
					{Key: "Authorization", Value: "Basic MY_BASIC_AUTH_TOKEN"},
					{Key: "Content-Type", Value: "MY_CONTENT_TYPE"},
				},
				Payload: []webhookintegration.Payload{
					{Key: "name", Value: "ali"},
					{Key: "birth_day", Value: "2020-12-12"},
					{Key: "mail", Value: "ali@mail.com"},
				},
				Method: webhookintegration.GETWebhookMethod,
				Url:    "https://eoc0z7vqfxu6io.m.pipedream.net",
			},
		},
		MessageID:         "13",
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
		log.Panicf("Error: %s", err)
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

	logger.L().Debug("Publish new processed event.")
}
