package managerinternalevent

import (
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/internalevent"
)

type SourceCreatedEventHandler struct {
	internalevent.Handler
}

func NewSourceCreatedEventHandler(adapter channel.Adapter, register internalevent.RegisterEventChannel) (SourceCreatedEventHandler, error) {
	handler := internalevent.New(adapter)
	eventHandler := SourceCreatedEventHandler{
		handler,
	}
	err := register(eventHandler.GetEventName())

	return eventHandler, err
}

func (m SourceCreatedEventHandler) Publish(msg *source.Source) error {
	return m.Handle(msg, m.GetEventName())
}

func (m SourceCreatedEventHandler) Consume() (<-chan channel.Message, error) {
	return m.GetOutputChannel(m.GetEventName())
}

func (m SourceCreatedEventHandler) GetEventName() string {
	return internalevent.EventNameSourceCreated
}
