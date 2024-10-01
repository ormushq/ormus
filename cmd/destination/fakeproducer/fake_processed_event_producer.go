package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"log/slog"
	"math/big"

	"os"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/contract/go/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/metricname"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	// generate fake processedEvent
	fakeEvent := &event.ProcessedEvent{
		SourceId:      "source-11",
		TracerCarrier: nil,
		Integration: &event.Integration{
			Id:       "sjdfkjsdkfjkdsfj",
			SourceId: "source-11",
			Name:     "Fake Integration",
			Metadata: &event.DestinationMetadata{
				Id:   "metadata-11",
				Name: "Test Metadata",
				Slug: event.DestinationType_webhook,
				Categories: []event.DestinationCategory{
					event.DestinationCategory_ANALYTICS,
					event.DestinationCategory_EMAIL_MARKETING,
				},
			},
			ConnectionType: event.ConnectionType_EVENT_STREAM,
			Enabled:        true,
			Config: &event.Integration_Webhook{
				Webhook: &event.WebhookConfig{
					Headers: map[string]string{
						"Authorization": "Basic MY_BASIC_AUTH_TOKEN",
						"Content-Type":  "MY_CONTENT_TYPE",
					},
					Payload: map[string]string{
						"name":      "ali",
						"birth_day": "2020-12-12",
						"mail":      "ali@mail.com",
					},
					Method: event.WebhookMethod_POST,
					Url:    "https://eoc0z7vqfxu6io.m.pipedream.net",
				},
			},
			CreatedAt: timestamppb.New(time.Now().Add(-1 * time.Hour)),
		},
		MessageId:         "11",
		EventType:         event.Type_TRACK,
		Version:           1,
		SentAt:            timestamppb.New(time.Now()),
		ReceivedAt:        timestamppb.New(time.Now()),
		OriginalTimestamp: timestamppb.New(time.Now()),
		Timestamp:         timestamppb.New(time.Now()),
		UserId:            "",
		AnonymousId:       "",
		Event:             "",
		Name:              "",
		GroupId:           "",
		PreviousId:        "",
		Context:           nil,
		Properties:        nil,
		Traits:            nil,
	}

	args := os.Args
	if len(args) > 1 && args[1] == "bulk" {
		for {
			publishEvent(fakeEvent, ch)
			l.Debug("Publish new processed event.")
			time.Sleep(time.Second)
		}
	} else {
		publishEvent(fakeEvent, ch)
		l.Debug("Publish new processed event.")
	}

	span.AddEvent("message-published")
	span.End()

	time.Sleep(time.Second * timeoutSeconds)
	close(done)
	wg.Wait()
}

func publishEvent(pe *event.ProcessedEvent, ch *amqp.Channel) {
	tracer := otela.NewTracer("FakerTracer")

	ctx, taskSpan := tracer.Start(context.Background(), "FakerTracer@publishEvent")

	pe.TracerCarrier = otela.GetCarrierFromContext(ctx)
	l := 3
	pe.MessageId = randSeq(l)
	pe.Integration.Id = randSeq(l)

	taskSpan.AddEvent("json-event-created", trace.WithAttributes(
		attribute.String("source-id", pe.SourceId),
		// TODO: write id service for processedevent
		attribute.String("event-id", ID(pe)),
		attribute.String("event-type", string(pe.EventType)),
	))

	taskSpan.End()

	meter := otel.Meter("FakerTracer@publishEvent")

	ppe, err := proto.Marshal(pe)

	err = ch.PublishWithContext(ctx,
		"processed-events-exchange", // exchange
		"pe.webhook",                // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        ppe,
		})
	failOnError(err, "Failed to publish a message")
	otela.IncrementFloat64Counter(ctx, meter, metricname.ProcessFlowOutputCore, "event_sent_to_destination")
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

func ID(e *event.ProcessedEvent) string {
	return e.SourceId + "-" + e.TracerCarrier["integration_id"]
}
