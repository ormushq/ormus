package eventhandler

import (
	"context"
	"fmt"
	"sync"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/ormushq/ormus/source/service/writekey"
)

type ProcessMessage func(ctx context.Context, msg channel.Message) error

type Consumer struct {
	broker          channel.Adapter
	writeKeyService writekey.Service
}

func New(broker channel.Adapter, writeKeyService writekey.Service) *Consumer {
	return &Consumer{
		broker:          broker,
		writeKeyService: writeKeyService,
	}
}

func (c Consumer) ProcessNewSourceEvent(ctx context.Context, msg channel.Message) error {
	decodedEvent := encoder.DecodeNewSourceEvent(string(msg.Body))
	// Log retrieval
	logger.L().Info(fmt.Sprintf("project id : %s, write key: %s, owner id: %s: has been retrieved",
		decodedEvent.ProjectId, decodedEvent.WriteKey, decodedEvent.OwnerId))

	err := c.writeKeyService.CreateNewWriteKey(ctx, decodedEvent.OwnerId, decodedEvent.ProjectId, decodedEvent.WriteKey)
	if err != nil {
		logger.L().Error("err on creating writekey in redis", "err msg:", err.Error())
		// TODO support Nack in pkg
	}

	logger.L().Debug("the message has been received")
	err = msg.Ack()
	if err != nil {
		logger.L().Debug("ack failed for message", "err msg:", err.Error())
	}

	return nil
}

func (c Consumer) Consume(ctx context.Context, queueName string, done <-chan bool, wg *sync.WaitGroup, processMessage ProcessMessage) {
	logger.L().Debug("Consumer started")
	wg.Add(1)
	go func() {
		defer wg.Done()

		msgChan, err := c.broker.GetOutputChannel(queueName)
		if err != nil {
			logger.L().Debug("error while subscribing to source topic")
		}
		for {
			select {
			case msg := <-msgChan:
				go func() {
					if err := processMessage(ctx, msg); err != nil {
						logger.L().Debug("error processing message", "err msg:", err.Error())
					}
				}()
			case <-done:
				return
			}
		}
	}()
}
