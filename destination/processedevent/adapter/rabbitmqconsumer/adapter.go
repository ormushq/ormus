package rabbitmqconsumer

import (
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
}

func New(c dconfig.RabbitMQConsumerConnection, topic dconfig.ConsumerTopic) *Consumer {
	return &Consumer{
		connectionConfig: c,
		topic:            topic,
	}
}

func (c *Consumer) Consume(done <-chan bool, wg *sync.WaitGroup) (<-chan event.ProcessedEvent, error) {
	// Start a goroutine to handle incoming messages
	// todo get size of channel from configs.
	channelSize := 100
	events := make(chan event.ProcessedEvent, channelSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(events)

		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", c.connectionConfig.User, c.connectionConfig.Password, c.connectionConfig.Host, c.connectionConfig.Port))
		failOnError(err, "Failed to connect to RabbitMQ")
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

		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = ch.QueueBind(
			q.Name,                      // queue name
			string(c.topic),             // routing key
			"processed-events-exchange", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")

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

		for {
			select {
			case msg := <-msgs:
				if len(msg.Body) == 0 {
					break
				}
				// Acknowledge the message
				err := msg.Ack(false)
				if err != nil {
					slog.Error(fmt.Sprintf("Failed to acknowledge message: %v", err))
				}

				e, uErr := taskentity.UnmarshalBytesToProcessedEvent(msg.Body)
				if uErr != nil {
					slog.Error(fmt.Sprintf("Failed to convert bytes to processed events: %v", uErr))

					break
				}

				events <- e

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
