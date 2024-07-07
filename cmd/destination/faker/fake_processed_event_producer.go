package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
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
	wg := &sync.WaitGroup{}
	done := make(chan bool)

	//----------------- Setup Logger -----------------//

	fileMaxSizeInMB := 10
	fileMaxAgeInDays := 30

	cfg := logger.Config{
		FilePath:         "./destination/logs.json",
		UseLocalTime:     false,
		FileMaxSizeInMB:  fileMaxSizeInMB,
		FileMaxAgeInDays: fileMaxAgeInDays,
	}

	logLevel := slog.LevelDebug

	opt := slog.HandlerOptions{
		// todo should level debug be read from config?
		Level: logLevel,
	}
	l := logger.New(cfg, &opt)

	// use slog as default logger.
	slog.SetDefault(l)

	//----------------- Setup Tracer -----------------//
	otelcfg := otela.Config{
		Endpoint:           config.C().Destination.Otel.Endpoint,
		ServiceName:        config.C().Destination.Otel.ServiceName + "/FakerEventGenerator",
		EnableMetricExpose: false,
	}

	err := otela.Configure(wg, done, otelcfg)
	failOnError(err, "Failed to configure otel")

	tracer := otela.NewTracer("FakerTracer")
	_, span := tracer.Start(context.Background(), "FakerTracer@FakerEventPreparation")

	rmqConsumerConnConfig := config.C().Destination.RabbitMQConsumerConnection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rmqConsumerConnConfig.User,
		rmqConsumerConnConfig.Password, rmqConsumerConnConfig.Host, rmqConsumerConnConfig.Port))
	span.AddEvent("rabbitmq-connection-established")

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
	span.AddEvent("channel-opened")

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
	span.AddEvent("exchange-declared")

	fakeIntegration := entity.Integration{
		ID:       "10",
		SourceID: "1",
		Metadata: entity.DestinationMetadata{
			ID:   "1",
			Name: "webhook",
			Slug: "webhook",
		},
	}

	ctx, taskSpan := tracer.Start(context.Background(), "FakerTracer@FakerEventPublisher")
	// generate fake processedEvent
	pageName := "Home"
	pe := event.ProcessedEvent{
		SourceID:          "1",
		TracerCarrier:     otela.GetCarrierFromContext(ctx),
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
	taskSpan.AddEvent("json-event-created", trace.WithAttributes(
		attribute.String("source-id", pe.SourceID),
		attribute.String("event-id", pe.ID()),
		attribute.String("event-type", string(pe.EventType)),
	))

	span.AddEvent("message-published")

	taskSpan.End()
	span.End()

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

	l.Debug("Publish new processed event.")

	time.Sleep(time.Second * 5)
	close(done)
	wg.Wait()
}
