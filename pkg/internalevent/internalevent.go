package internalevent

import (
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

func (h Handler) Handle(msg proto.Message, channelName string) error {
	inputChannel, err := h.adapter.GetInputChannel(channelName)
	if err != nil {
		return err
	}
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	inputChannel <- body

	return nil
}

func (h Handler) GetOutputChannel(channelName string) (<-chan channel.Message, error) {
	return h.adapter.GetOutputChannel(channelName)
}
