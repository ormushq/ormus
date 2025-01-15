package scheduler

import (
	"context"
	"fmt"
	"sync"

	proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/eventhandler"
)

type Repository interface {
	GetAllUnDeliveredEvents() ([]*proto.NewEvent, error)
}

type Scheduler struct {
	repo         Repository
	eventAdapter eventhandler.Publisher
	config       source.Config
	wg           *sync.WaitGroup
}

func New(repo Repository, eventAdapter eventhandler.Publisher, config source.Config, wg *sync.WaitGroup) Scheduler {
	return Scheduler{
		repo:         repo,
		eventAdapter: eventAdapter,
		config:       config,
		wg:           wg,
	}
}

func (s Scheduler) PublishUndeliveredEvents(ctx context.Context) error {
	logger.L().Info("scheduler started ...")
	select {
	case <-ctx.Done():
		logger.L().Error("the context has been canceled", "reason", "context cancelled")

		return nil
	default:
		events, err := s.repo.GetAllUnDeliveredEvents()
		if err != nil {
			logger.L().Error("unable to get all undelivered events", "error", err.Error())

			return err
		}
		err = s.eventAdapter.Publish(ctx, s.config.NewEventQueueName, s.wg, events)
		if err != nil {
			logger.L().Error("unable to publish undelivered events", "error", err.Error())

			return err
		}
		logger.L().Info(fmt.Sprintf("%d undelivered events successfully published", len(events)))
	}

	return nil
}
