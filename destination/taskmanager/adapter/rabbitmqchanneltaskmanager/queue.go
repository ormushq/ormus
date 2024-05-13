package rabbitmqchanneltaskmanager

import (
	"encoding/json"
	"github.com/ormushq/ormus/destination/channel"
	rbbitmqchannel "github.com/ormushq/ormus/destination/channel/adapter/rabbitmq"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/event"
	"log"
	"log/slog"
	"sync"
)

type Queue struct {
	channel chan<- []byte
	name    string
	config  dconfig.RabbitMQTaskManagerConnection
}

const timeoutSeconds = 5

func newQueue(done <-chan bool, wg *sync.WaitGroup, c dconfig.RabbitMQTaskManagerConnection, n string, reconnectSecond int) *Queue {
	inputChannelAdapter := rbbitmqchannel.New(done, wg, dconfig.RabbitMQConsumerConnection{
		User:            c.User,
		Password:        c.Password,
		Host:            c.Host,
		Port:            c.Port,
		Vhost:           "/",
		ReconnectSecond: reconnectSecond,
	})

	inputChannelAdapter.NewChannel(n, channel.InputOnlyMode, 0, 1, 5)
	chanel, err := inputChannelAdapter.GetInputChannel(n)
	if err != nil {
		log.Fatal(err)
	}
	return &Queue{
		name:    n,
		config:  c,
		channel: chanel,
	}
}

func (q *Queue) Enqueue(pe event.ProcessedEvent) error {
	// Convert Processed event to json
	jsonEvent, err := json.Marshal(pe)
	if err != nil {
		slog.Error("Error:", err)

		return err
	}
	q.channel <- jsonEvent

	return nil
}
