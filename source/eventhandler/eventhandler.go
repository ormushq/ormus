package eventhandler

import (
	"context"
	"errors"
	"fmt"
	"sync"

	proto_source "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/service/event"
	"github.com/ormushq/ormus/source/service/writekey"
)

type ProcessMessage func(ctx context.Context, msg channel.Message) error

type Consumer struct {
	broker          channel.Adapter
	writeKeyService writekey.Service
	eventService    event.Service
	retryNumber     int
}

func NewConsumer(broker channel.Adapter, writeKeyService writekey.Service, eventService event.Service, retryNumber int) *Consumer {
	return &Consumer{
		broker:          broker,
		writeKeyService: writeKeyService,
		eventService:    eventService,
		retryNumber:     retryNumber,
	}
}

func (c Consumer) Consume(ctx context.Context, queueName string, done <-chan bool, wg *sync.WaitGroup, processMessage ProcessMessage) {
	logger.L().Debug("Consumer started")
	wg.Add(1)

	go func() {
		defer wg.Done()

		msgChan, err := c.broker.GetOutputChannel(queueName)
		if err != nil {
			logger.L().Error(fmt.Sprintf("error while subscribing to source topic %s", err.Error()))
		}
		for {
			select {
			case msg := <-msgChan:
				go func() {
					if err := processMessage(ctx, msg); err != nil {
						logger.L().Debug("error processing message", "error", err.Error())
					}
				}()
			case <-done:
				return
			}
		}
	}()
}

type Publisher struct {
	broker channel.Adapter
}

func NewPublisher(broker channel.Adapter) Publisher {
	return Publisher{
		broker: broker,
	}
}

func (p Publisher) Publish(ctx context.Context, queueName string, wg *sync.WaitGroup, messages []*proto_source.NewEvent) error {
	msgChan, err := p.broker.GetInputChannel(queueName)
	if err != nil {
		logger.L().Error("error while subscribing to source topic")

		return richerror.New("source.event_handler").
			WithKind(richerror.KindUnexpected).
			WithWrappedError(err)
	}
	wg.Add(len(messages))
	for _, message := range messages {
		go func(message *proto_source.NewEvent) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				if errors.Is(ctx.Err(), context.DeadlineExceeded) {
					logger.L().Error("context timeout exceeded")
				} else if errors.Is(ctx.Err(), context.Canceled) {
					logger.L().Error("context cancelled by force. whole process is complete")
				}

				return
			case msgChan <- []byte(encoder.EncodeNewEvent(message)):
			}
		}(message)
	}

	return nil
}
