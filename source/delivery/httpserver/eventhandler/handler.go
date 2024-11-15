package eventhandler

import (
	"github.com/ormushq/ormus/source/service/event"
	"github.com/ormushq/ormus/source/validator/eventvalidator/eventvalidator"
)

type Handler struct {
	eventSvc       event.Service
	eventValidator eventvalidator.Validator
}

func New(eventSvc event.Service, eventValidator eventvalidator.Validator) Handler {
	return Handler{
		eventSvc:       eventSvc,
		eventValidator: eventValidator,
	}
}
