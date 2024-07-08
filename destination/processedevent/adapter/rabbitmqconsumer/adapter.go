package rabbitmqconsumer

import (
	"context"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/pkg/metricname"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	//"context"
	"fmt"
	"log"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	topic            dconfig.ConsumerTopic
	connectionConfig dconfig.RabbitMQConsumerConnection
	ctx              context.Context
}

func New(ctx context.Context, c dconfig.RabbitMQConsumerConnection, topic dconfig.ConsumerTopic) *Consumer {
	return &Consumer{
		ctx:              ctx,
		connectionConfig: c,
		topic:            topic,
	}
}

func (c *Consumer) Consume(done <-chan bool, wg *sync.WaitGroup) (<-chan event.ProcessedEvent, error) {
	tracer := otela.NewTracer("rabbitmqconsumer")

	// Start a goroutine to handle incoming messages
	// todo get size of channel from configs.
	channelSize := 100
	events := make(chan event.ProcessedEvent, channelSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(events)
		_, span := tracer.Start(c.ctx, "rabbitmqconsumer@consume")

		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", c.connectionConfig.User, c.connectionConfig.Password, c.connectionConfig.Host, c.connectionConfig.Port))
		failOnError(err, "Failed to connect to RabbitMQ")
		defer func(conn *amqp.Connection) {
			err = conn.Close()
			failOnError(err, "Failed to close RabbitMQ connection")
		}(conn)
		span.AddEvent("connection-established")

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer func(ch *amqp.Channel) {
			err = ch.Close()
			failOnError(err, "Failed to close channel")
		}(ch)
		span.AddEvent("channel-established")

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
		span.AddEvent("exchange-declare")

		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")
		span.AddEvent("queue-declare")

		err = ch.QueueBind(
			q.Name,                      // queue name
			string(c.topic),             // routing key
			"processed-events-exchange", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
		span.AddEvent("queue-bind")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // arguments
		)
		failOnError(err, "failed to consume")
		span.AddEvent("consume-started")
		span.End()

		meter := otel.GetMeterProvider().Meter("rabbitmqconsumer@Consume")

		for {
			select {
			case msg := <-msgs:
				wg.Add(1)
				go func() {
					defer wg.Done()
					if len(msg.Body) == 0 {
						return
					}

					e, uErr := taskentity.UnmarshalBytesToProcessedEvent(msg.Body)
					if uErr != nil {
						slog.Error(fmt.Sprintf("Failed to convert bytes to processed events: %v", uErr))
						otela.IncrementFloat64Counter(context.Background(), meter, metricname.DESTINATION_INPUT_UNMARSHAL_ERROR, "processed_event_unmarshal_error")

						return
					}
					ctx := otela.GetContextFromCarrier(e.TracerCarrier)
					ctx, span = tracer.Start(ctx, "rabbitmqconsumer@StartProccessEvent")
					defer span.End()

					otela.IncrementFloat64Counter(ctx, meter, metricname.PROCESS_FLOW_INPUT_DESTINATION, "event_received_in_destination")

					span.AddEvent("process-started")

					e.TracerCarrier = otela.GetCarrierFromContext(ctx)
					events <- e

					// Acknowledge the message
					err := msg.Ack(false)
					if err != nil {
						otela.IncrementFloat64Counter(ctx, meter, metricname.DESTINATION_INPUT_ACK_ERROR, "processed_event_ack_error")

						span.AddEvent("error-on-ack", trace.WithAttributes(
							attribute.String("error", err.Error())))

						slog.Error(fmt.Sprintf("Failed to acknowledge message: %v", err))
					}
				}()

			case <-done:

				return
			}
		}
	}()

	return events, nil
}

func (c *Consumer) Close() error {
	// todo close rabbit consumer
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
