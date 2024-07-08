package rabbitmqchanneltaskmanager

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/pkg/metricname"
	"go.opentelemetry.io/otel"
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
	meter := otel.Meter("rabbitmqchanneltaskmanager@Consume")

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-c.messageChannel:
				wg.Add(1)
				go func() {
					defer wg.Done()
					otela.IncrementFloat64Counter(context.Background(), meter, metricname.PROCESS_FLOW_INPUT_DESTINATION_WORKER, "event_received_in_worker")

					e, err := taskentity.UnmarshalBytesToProcessedEvent(msg.Body)
					if err != nil {
						otela.IncrementFloat64Counter(context.Background(), meter, metricname.DESTINATION_WORKER_INPUT_UNMARSHAL_ERROR, "process_event_unmarshal_error")

						printWorkersError(err, "Failed to unmarshall message")

						return
					}
					tracer := otela.NewTracer("rabbitmqchanneltaskmanager")
					ctx, span := tracer.Start(otela.GetContextFromCarrier(e.TracerCarrier), "rabbitmqchanneltaskmanager@TaskConsumer")
					defer span.End()
					e.TracerCarrier = otela.GetCarrierFromContext(ctx)
					span.AddEvent("process-event-consumed")

					eventsChannel <- e
					aErr := msg.Ack()
					if aErr != nil {
						otela.IncrementFloat64Counter(ctx, meter, metricname.DESTINATION_WORKER_INPUT_ACK_ERROR, "process_event_ack_error")

						printWorkersError(aErr, "Failed to acknowledge message")
						span.AddEvent("error-on-ack", trace.WithAttributes(
							attribute.String("error", err.Error())))

						return
					}

					otela.IncrementFloat64Counter(ctx, meter, metricname.DESTINATION_WORKER_EVENT_SEND_TO_WORKER, "process_event_publish_to_event_channel")
					span.AddEvent("process-event-publish-to-event-channel")
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
