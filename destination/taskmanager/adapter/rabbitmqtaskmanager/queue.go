package rabbitmqtaskmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	name   string
	config dconfig.RabbitMQTaskManagerConnection
}

const timeoutSeconds = 5

func newQueue(c dconfig.RabbitMQTaskManagerConnection, n string) *Queue {
	return &Queue{
		name:   n,
		config: c,
	}
}

func (q *Queue) Enqueue(pe event.ProcessedEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	// Connect to RabbitMQ
	connectionConfig := q.config
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", connectionConfig.User, connectionConfig.Password, connectionConfig.Host, connectionConfig.Port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare queue
	rq, err := ch.QueueDeclare(
		q.name, // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Convert Processed event to json
	jsonEvent, err := json.Marshal(pe)
	if err != nil {
		fmt.Println("Error:", err)

		return err
	}

	// Publish message to queue
	err = ch.PublishWithContext(ctx,
		"",      // exchange
		rq.Name, // routing key
		false,   // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         jsonEvent,
		})

	failOnError(err, "Failed to publish a message")

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
