package eventmanager

import (
	"github.com/ormushq/ormus/contract/go/internalevent"
)

type EventMessage struct {
	Event *internalevent.Event
	Ack   func() error
}
type CreateChannelFunc func(channelName string) error

func NewEventMessage(ev *internalevent.Event, ack func() error) EventMessage {
	return EventMessage{
		Event: ev,
		Ack:   ack,
	}
}
