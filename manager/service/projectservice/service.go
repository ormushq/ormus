package projectservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/channel"
)

type Repository interface {
	Create(name, ID string) (entity.Project, error)
}

type Service struct {
	repo           Repository
	internalBroker channel.Adapter
}

func New(repository Repository, internalBroker channel.Adapter) *Service {
	return &Service{
		repo:           repository,
		internalBroker: internalBroker,
	}
}
