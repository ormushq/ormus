package writekeyadapter

import (
	"context"
	"encoding/json"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"github.com/ormushq/ormus/source/params"
	"github.com/ormushq/ormus/source/service/writekey"
	"sync"
)

type Consumer struct {
	internalBroker  rbbitmqchannel.ChannelAdapter
	WriteKeyService writekey.Service
}

func New(internalBroker rbbitmqchannel.ChannelAdapter, WriteKeyService writekey.Service) *Consumer {
	return &Consumer{
		internalBroker:  internalBroker,
		WriteKeyService: WriteKeyService,
	}
}
func (c Consumer) ConsumeWriteKey(ctx context.Context, done <-chan bool, wg *sync.WaitGroup) {
	logger.L().Debug("Write Key Consumer started")
	internalBroker, err := c.internalBroker.GetOutputChannel("new-source-created")
	if err != nil {
		logger.L().Debug("error on getting internal broker channel")
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
		case msg := <-internalBroker:
			var req params.WriteKey
			err := json.Unmarshal(msg.Body, &req)
			if err != nil {
				logger.L().Debug("error on unmarshalling json")
			}
			err = c.WriteKeyService.CreateNewWriteKey(ctx, req.OwnerID, req.ProjectID, req.WriteKey)
			if err != nil {
				logger.L().Error("err on creating writekey in redis", "err msg:", err.Error())
				break
			}
			err = msg.Ack()
			if err != nil {
				logger.L().Error("err on acking message", "err msg:", err)
				break
			}

		case <-done:
			return
		}
	}()

}
