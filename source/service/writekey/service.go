package writekey

import (
	"context"
	"github.com/ormushq/ormus/manager/entity"
)

type Repository interface {
	IsValidWriteKey(ctx context.Context, writeKey string) (*entity.WriteKeyMetaData, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
