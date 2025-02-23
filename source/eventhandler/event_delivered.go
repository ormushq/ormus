package eventhandler

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/ormushq/ormus/pkg/retry"
)

func (c Consumer) EventHasDeliveredToDestination(ctx context.Context, msg channel.Message) error {
	decodedEvent := encoder.DecodeProcessedEvent(string(msg.Body))
	err := c.eventService.EventHasDelivered(ctx, decodedEvent)
	if err != nil {
		logger.L().Error(fmt.Sprintf("err on change delivered status of events : %s", err.Error()))
		fn := func() error {
			return c.eventService.EventHasDelivered(ctx, decodedEvent)
		}
		err = retry.Do(fn, c.retryNumber)
		if err != nil {
			logger.L().Error(err.Error())
		}
	}
	logger.L().Info(fmt.Sprintf("processed event event array: %v has been retrieved", decodedEvent.Events))
	err = msg.Ack()
	if err != nil {
		logger.L().Debug(fmt.Sprintf("ack failed for message : %s", err.Error()))
		return err
	}
	return nil
}
