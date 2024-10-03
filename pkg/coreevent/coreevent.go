package coreevent

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/pkg/channel"
	"google.golang.org/protobuf/proto"
)

type RegisterEventChannel func(channelName string) error

type Handler struct {
	adapter channel.Adapter
}

func New(adapter channel.Adapter) Handler {
	return Handler{
		adapter: adapter,
	}
}

func (h Handler) Handle(ctx context.Context, msg proto.Message, channelName string) error {
	inputChannel := h.adapter.GetInputChannel(channelName)
	if inputChannel == nil {
		return fmt.Errorf("no input channel found for event %s", channelName)
	}
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	inputChannel <- channel.BrokerMessage{
		Ctx:  ctx,
		Body: body,
	}

	return nil
}

func (h Handler) GetOutputChannel(channelName string, exchanges ...string) (<-chan channel.BrokerMessage, error) {
	return h.adapter.GetOutputChannel(channelName, exchanges...)
}
