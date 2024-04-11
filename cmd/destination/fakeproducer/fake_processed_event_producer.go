package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	"github.com/ormushq/ormus/pkg/metricname"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
		Exporter:           otela.ExporterGrpc,
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

	_ = entity.Integration{
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
				Method: webhookintegration.POSTWebhookMethod,
				URL:    "https://eoc0z7vqfxu6io.m.pipedream.net",
			},
		},
		MessageID:         "14",
		EventType:         "page",
		Name:              &pageName,
		Version:           1,
		SentAt:            time.Now(),
		ReceivedAt:        time.Now(),
		OriginalTimestamp: time.Now(),
		Timestamp:         time.Now(),
	}

	args := os.Args
	if len(args) > 1 && args[1] == "bulk" {
		for {
			publishEvent(pe, ch)
			l.Debug("Publish new processed event.")
			time.Sleep(time.Second)
		}
	} else {
		publishEvent(pe, ch)
		l.Debug("Publish new processed event.")
	}

	span.AddEvent("message-published")
	span.End()

	time.Sleep(time.Second * timeoutSeconds)
	close(done)
	wg.Wait()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}

		b[i] = letters[nBig.Int64()]
	}

	return string(b)
}

func publishEvent(pe event.ProcessedEvent, ch *amqp.Channel) {
	tracer := otela.NewTracer("FakerTracer")

	ctx, taskSpan := tracer.Start(context.Background(), "FakerTracer@publishEvent")

	pe.TracerCarrier = otela.GetCarrierFromContext(ctx)
	l := 3
	pe.MessageID = randSeq(l)
	pe.Integration.ID = randSeq(l)

	jpe, err := json.Marshal(pe)
	if err != nil {
		log.Panicf("Error: %s", err)
	}
	taskSpan.AddEvent("json-event-created", trace.WithAttributes(
		attribute.String("source-id", pe.SourceID),
		attribute.String("event-id", pe.ID()),
		attribute.String("event-type", string(pe.EventType)),
	))

	taskSpan.End()

	meter := otel.Meter("FakerTracer@publishEvent")
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
	otela.IncrementFloat64Counter(ctx, meter, metricname.ProcessFlowOutputCore, "event_sent_to_destination")
}
