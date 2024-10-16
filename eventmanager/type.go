package eventmanager

import (
	"github.com/ormushq/ormus/contract/go/internalevent"
)

type EventMessage struct {
	Event *internalevent.Event
	Ack   func() error
}

func NewEventMessage(ev *internalevent.Event, ack func() error) EventMessage {
	return EventMessage{
		Event: ev,
		Ack:   ack,
	}
}
