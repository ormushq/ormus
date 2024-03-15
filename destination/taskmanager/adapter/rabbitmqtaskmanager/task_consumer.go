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
	channelSize := 6
	eventsChannel := make(chan event.ProcessedEvent, channelSize)
	wg.Add(1)

	go func() {
		defer wg.Done()

		connectionConfig := c.ConnectionConfig
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", connectionConfig.User, connectionConfig.Password, connectionConfig.Host, connectionConfig.Port))
		panicOnWorkersError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		panicOnWorkersError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			c.QueueName, // name
			true,        // durable
			false,       // delete when unused
			false,       // exclusive
			false,       // no-wait
			nil,         // arguments
		)
		panicOnWorkersError(err, "Failed to declare a queue")

		err = ch.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
		)
		panicOnWorkersError(err, "Failed to set QoS")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
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
