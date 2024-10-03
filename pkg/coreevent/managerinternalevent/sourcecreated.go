package managerinternalevent

import (
	"context"

	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/coreevent"
)

type SourceCreatedEventHandler struct {
	coreevent.Handler
}

func NewSourceCreatedEventHandler(adapter channel.Adapter, register coreevent.RegisterEventChannel) (SourceCreatedEventHandler, error) {
	handler := coreevent.New(adapter)
	eventHandler := SourceCreatedEventHandler{
		handler,
	}
	err := register(eventHandler.GetEventName())

	return eventHandler, err
}

func (m SourceCreatedEventHandler) Publish(ctx context.Context, msg *source.MessageDeliveryStatus) error {
	return m.Handle(ctx, msg, m.GetEventName())
}

func (m SourceCreatedEventHandler) Consume(exchanges ...string) (<-chan channel.BrokerMessage, error) {
	return m.GetOutputChannel(m.GetEventName(), exchanges...)
}

func (m SourceCreatedEventHandler) GetEventName() string {
	return managementconst.EventMessageSetDeliveryStatus
}
