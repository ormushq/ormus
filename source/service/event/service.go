package event

import (
	"context"
	"sync"

	"github.com/ormushq/ormus/contract/go/destination"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/params"
)

type Repository interface {
	CreateNewEvent(ctx context.Context, evt []event.CoreEvent, wg *sync.WaitGroup, queueName string) ([]string, error)
	EventHasDelivered(ctx context.Context, evt *destination.DeliveredEventsList) error
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

func (s Service) CreateNewEvent(ctx context.Context, newEvent []params.TrackEventRequest, invalidWriteKeys []string) (*params.TrackEventResponse, error) {
	events := make([]event.CoreEvent, 0)
	for _, e := range newEvent {
		events = append(events, event.CoreEvent{
			Name:       e.Name,
			WriteKey:   e.WriteKey,
			Event:      e.Event,
			SendAt:     e.SendAt,
			ReceivedAt: e.ReceivedAt,
			Timestamp:  e.Timestamp,
			Type:       event.Type(e.Type),
			Properties: (*event.Properties)(&e.Properties),
		})
	}

	ids, err := s.eventRepo.CreateNewEvent(ctx, events, s.wg, s.config.NewEventQueueName)
	if err != nil {
		logger.L().Error(err.Error())

		return nil, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
	}

	return &params.TrackEventResponse{
		ID:               ids,
		InvalidWriteKeys: invalidWriteKeys,
		Success:          len(ids),
		FAIL:             len(invalidWriteKeys),
	}, nil
}

func (s Service) EventHasDelivered(ctx context.Context, evt *destination.DeliveredEventsList) error {
	err := s.eventRepo.EventHasDelivered(ctx, evt)
	if err != nil {
		logger.L().Error(err.Error())

		return richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
	}

	return nil
}
