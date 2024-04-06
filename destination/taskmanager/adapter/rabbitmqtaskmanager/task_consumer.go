package rabbitmqtaskmanager

import (
	"fmt"
	"sync"

	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	ConnectionConfig dconfig.RabbitMQTaskManagerConnection
	QueueName        string
}

func NewTaskConsumer(cnf dconfig.RabbitMQTaskManagerConnection, queueName string) Consumer {
	return Consumer{
		ConnectionConfig: cnf,
		QueueName:        queueName,
	}
}

func (c Consumer) Consume(done <-chan bool, wg *sync.WaitGroup) (<-chan event.ProcessedEvent, error) {
	// todo use configs for size of channel
	channelSize := 100
	eventsChannel := make(chan event.ProcessedEvent, channelSize)
	wg.Add(1)

	go func() {
		defer wg.Done()

		// Connect to rabbitMQ
		connectionConfig := c.ConnectionConfig
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", connectionConfig.User, connectionConfig.Password, connectionConfig.Host, connectionConfig.Port))
		panicOnWorkersError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		panicOnWorkersError(err, "Failed to open a channel")
		defer ch.Close()

		// Declare queue (create or check for)
		q, err := ch.QueueDeclare(
			c.QueueName, // Name of queue
			true,        // A durable queue persists on disk, meaning its messages won't be lost even if the RabbitMQ server restarts
			false,       // The queue should not be automatically deleted when it's no longer in use
			false,       // An exclusive queue can only be used by one connection/consumer at a time.
			false,       // no-wait
			nil,         // arguments
		)
		panicOnWorkersError(err, "Failed to declare a queue")

		err = ch.Qos(
			1,     // Maximum number of messages RabbitMQ will deliver to the consumer before waiting for acknowledgments.
			0,     // Maximum size (in bytes) of message content. 0 means no limit.
			false, // Apply settings only to this specific channel (not globally).
		)
		panicOnWorkersError(err, "Failed to set QoS")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // RabbitMQ will generate a unique consumer tag
			false,  // Determines whether messages should be automatically acknowledged
			false,  // Specifies only this consumer can access the queue or not
			false,  // Specifies whether the server should not deliver messages published by the same connection.
			false,  // no-wait
			nil,    // args
		)
		panicOnWorkersError(err, "Failed to register a consumer")

		for {
			select {
			case msg := <-msgs:
				aErr := msg.Ack(false)
				if aErr != nil {
					printWorkersError(err, "Failed to acknowledge message")

					break
				}
				e, err := taskentity.UnmarshalBytesToProcessedEvent(msg.Body)
				if err != nil {
					printWorkersError(err, "Failed to unmarshall message")

					break
				}
				eventsChannel <- e
			case <-done:

				return
			}
		}
	}()

	return eventsChannel, nil
}
