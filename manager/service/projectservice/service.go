package projectservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/validator/projectvalidator"
	"github.com/ormushq/ormus/pkg/channel"
)

type Repository interface {
	Create(project entity.Project) (entity.Project, error)
	GetWithID(ID string) (entity.Project, error)
	Update(project entity.Project) (entity.Project, error)
	Delete(project entity.Project) error
	List(userID string, lastToken int64, limit int) ([]entity.Project, error)
	HaseMore(userID string, lastToken int64, perPage int) (bool, error)
}

type Service struct {
	repo           Repository
	internalBroker channel.Adapter
	validator      projectvalidator.Validator
}

func New(repository Repository, internalBroker channel.Adapter, validator projectvalidator.Validator) Service {
	return Service{
		repo:           repository,
		internalBroker: internalBroker,
		validator:      validator,
	}
}
