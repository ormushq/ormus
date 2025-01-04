package event

import (
	"context"
	"sync"

	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/params"
)

type Repository interface {
	CreateNewEvent(ctx context.Context, evt event.CoreEvent, wg *sync.WaitGroup, queueName string) (string, error)
}

type Service struct {
	eventRepo Repository
	config    source.Config
	wg        *sync.WaitGroup
}

func New(eventRepo Repository, config source.Config, wg *sync.WaitGroup) *Service {
	return &Service{
		eventRepo: eventRepo,
		config:    config,
		wg:        wg,
	}
}

func (s Service) CreateNewEvent(ctx context.Context, newEvent params.TrackEventRequest) (*params.TrackEventResponse, error) {
	e := event.CoreEvent{
		Name:       newEvent.Name,
		WriteKey:   newEvent.WriteKey,
		Event:      newEvent.Event,
		SendAt:     newEvent.SendAt,
		ReceivedAt: newEvent.ReceivedAt,
		Timestamp:  newEvent.Timestamp,
		Type:       event.Type(newEvent.Type),
		Properties: (*event.Properties)(&newEvent.Properties),
	}

	id, err := s.eventRepo.CreateNewEvent(ctx, e, s.wg, s.config.NewEventQueueName)
	if err != nil {
		logger.L().Error(err.Error())

		return nil, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
	}

	return &params.TrackEventResponse{
		ID: id,
	}, nil
}
