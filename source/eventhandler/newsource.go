package eventhandler

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/ormushq/ormus/source/service/writekey"
	"sync"
)

type Consumer struct {
	broker          writekey.ConsumerRepo
	writeKeyService writekey.Service
}

func New(Broker writekey.ConsumerRepo, WriteKeyService writekey.Service) *Consumer {
	return &Consumer{
		broker:          Broker,
		writeKeyService: WriteKeyService,
	}
}
func (c Consumer) ConsumeWriteKey(ctx context.Context, queueName string, done <-chan bool, wg *sync.WaitGroup) {
	logger.L().Debug("Write Key Consumer started")
	wg.Add(1)
	go func() {
		defer wg.Done()
		msgchan, err := c.broker.Subscribe(queueName)
		if err != nil {
			logger.L().Debug("error while subscribing to source topic")
		}
		for {
			select {
			case msg := <-msgchan:
				decodedEvent := encoder.DecodeNewSourceEvent(string(msg.Body))
				logger.L().Info(fmt.Sprintf("project id : %s, write key: %s, owner id: %s:  has been retrieved",
					decodedEvent.ProjectId, decodedEvent.WriteKey, decodedEvent.OwnerId))
				err = c.writeKeyService.CreateNewWriteKey(ctx, decodedEvent.OwnerId, decodedEvent.ProjectId, decodedEvent.WriteKey)
				if err != nil {
					logger.L().Error("err on creating writekey in redis", "err msg:", err.Error())

					break
				}

				logger.L().Debug("the message has been received")
				err = c.broker.Ack(msg)
				if err != nil {
					logger.L().Debug("ack failed in source")
				}
			case <-done:
				return
			}
		}
	}()

}
