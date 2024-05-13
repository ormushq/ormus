package rabbitmqchanneltaskmanager

import (
	"github.com/ormushq/ormus/destination/channel"
	rbbitmqchannel "github.com/ormushq/ormus/destination/channel/adapter/rabbitmq"
	"log"
	"sync"

	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/event"
)

type Consumer struct {
	ConnectionConfig dconfig.RabbitMQTaskManagerConnection
	QueueName        string
	channelSize      int
	reconnectSecond  int
	numberInstant    int
	maxRetryPolicy   int
}

func NewTaskConsumer(cnf dconfig.RabbitMQTaskManagerConnection, queueName string,
	channelSize, reconnectSecond, numberInstant, maxRetryPolicy int) Consumer {
	return Consumer{
		ConnectionConfig: cnf,
		QueueName:        queueName,
		channelSize:      channelSize,
		reconnectSecond:  reconnectSecond,
		numberInstant:    numberInstant,
		maxRetryPolicy:   maxRetryPolicy,
	}
}

func (c Consumer) Consume(done <-chan bool, wg *sync.WaitGroup) (<-chan event.ProcessedEvent, error) {

	outputChannelAdapter := rbbitmqchannel.New(done, wg, dconfig.RabbitMQConsumerConnection{
		User:            c.ConnectionConfig.User,
		Password:        c.ConnectionConfig.Password,
		Host:            c.ConnectionConfig.Host,
		Port:            c.ConnectionConfig.Port,
		Vhost:           "/",
		ReconnectSecond: c.reconnectSecond,
	})

	outputChannelAdapter.NewChannel(c.QueueName, channel.OutputOnly, c.channelSize, c.numberInstant, c.maxRetryPolicy)

	eventsChannel := make(chan event.ProcessedEvent, c.channelSize)

	wg.Add(1)

	msgs, err := outputChannelAdapter.GetOutputChannel(c.QueueName)
	if err != nil {
		return nil, err
	}

	go func() {
		defer wg.Done()
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

func printWorkersError(err error, msg string) {
	log.Printf("%s: %s", msg, err)
}
