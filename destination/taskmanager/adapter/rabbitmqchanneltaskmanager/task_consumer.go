package rabbitmqchanneltaskmanager

import (
	"fmt"
	"sync"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
)

type Consumer struct {
	messageChannel <-chan channel.Message
	channelSize    int
}

func NewTaskConsumer(messageChannel <-chan channel.Message, channelSize int) Consumer {
	return Consumer{
		channelSize:    channelSize,
		messageChannel: messageChannel,
	}
}

func (c Consumer) Consume(done <-chan bool, wg *sync.WaitGroup) (<-chan event.ProcessedEvent, error) {
	eventsChannel := make(chan event.ProcessedEvent, c.channelSize)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-c.messageChannel:
				e, err := taskentity.UnmarshalBytesToProcessedEvent(msg.Body)
				if err != nil {
					printWorkersError(err, "Failed to unmarshall message")

					break
				}

				eventsChannel <- e
				aErr := msg.Ack()
				if aErr != nil {
					printWorkersError(aErr, "Failed to acknowledge message")

					break
				}
			case <-done:

				return
			}
		}
	}()

	return eventsChannel, nil
}

func printWorkersError(err error, msg string) {
	logger.L().Error(fmt.Sprintf("%s: %s", msg, err))
}
