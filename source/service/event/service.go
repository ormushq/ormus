package event

import (
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/params"
)

type Repository interface {
	CreateNewEvent(evt event.CoreEvent) (string, error)
}

type Service struct {
	eventRepo Repository
}

func New(eventRepo Repository) *Service {
	return &Service{eventRepo: eventRepo}
}

func (s Service) CreateNewEvent(newEvent params.TrackEventRequest) (*params.TrackEventResponse, error) {
	e := event.CoreEvent{
		Name:       newEvent.Name,
		WriteKey:   newEvent.Event,
		Event:      newEvent.Event,
		SendAt:     newEvent.SendAt,
		ReceivedAt: newEvent.ReceivedAt,
		Timestamp:  newEvent.Timestamp,
		Type:       event.Type(newEvent.Type),
		Properties: (*event.Properties)(&newEvent.Properties),
	}

	id, err := s.eventRepo.CreateNewEvent(e)
	if err != nil {
		return nil, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
	}

	return &params.TrackEventResponse{
		ID: id,
	}, nil
}
