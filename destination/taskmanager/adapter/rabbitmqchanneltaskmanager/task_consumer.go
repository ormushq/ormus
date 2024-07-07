package rabbitmqchanneltaskmanager

import (
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
				wg.Add(1)
				go func() {
					defer wg.Done()
					e, err := taskentity.UnmarshalBytesToProcessedEvent(msg.Body)
					if err != nil {
						printWorkersError(err, "Failed to unmarshall message")

						return
					}
					tracer := otela.NewTracer("rabbitmqchanneltaskmanager")
					ctx, span := tracer.Start(otela.GetContextFromCarrier(e.TracerCarrier), "rabbitmqchanneltaskmanager@TaskConsumer")
					defer span.End()
					e.TracerCarrier = otela.GetCarrierFromContext(ctx)
					span.AddEvent("task-consumed")

					eventsChannel <- e
					aErr := msg.Ack()
					if aErr != nil {
						printWorkersError(aErr, "Failed to acknowledge message")
						span.AddEvent("error-on-ack", trace.WithAttributes(
							attribute.String("error", err.Error())))

						return
					}
					span.AddEvent("task-send-to-event-channel")
				}()
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
