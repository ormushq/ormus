package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/ormushq/ormus/logger"
)

type Task func(ctx context.Context) error

type Scheduler struct {
	done chan bool
	wg   *sync.WaitGroup
	sch  *gocron.Scheduler
}

func New(done chan bool, wg *sync.WaitGroup) Scheduler {
	return Scheduler{
		done: done,
		wg:   wg,
		sch:  gocron.NewScheduler(time.UTC),
	}
}

func (s Scheduler) Start(ctx context.Context, period int, task Task) {
	_, err := s.sch.Every(period).Minute().Do(task, ctx)
	if err != nil {
		logger.L().Error("scheduler had problem in publishing the undelivered events", "error", err.Error())
	}
	s.sch.StartAsync()
	defer s.wg.Done()
	<-ctx.Done()
	logger.L().Info("the context has been canceled")
	close(s.done)
	s.sch.Stop()
}
