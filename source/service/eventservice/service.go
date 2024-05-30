package eventservice

import (
	"context"
	"github.com/ormushq/ormus/event"
)

type EventRepo interface {
	InsertEvent(ctx context.Context, event event.CoreEvent) (event.CoreEvent, error)
}

type Service struct {
	repo EventRepo
}

func New(repo EventRepo) Service {
	return Service{
		repo: repo,
	}
}
