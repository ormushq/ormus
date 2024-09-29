package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/validator/sourcevalidator"
)

type SourceRepo interface {
	Create(source entity.Source) (entity.Source, error)
	Update(id string, source *entity.Source) (*managerparam.UpdateSourceResponse, error)
	Delete(id, userID string) error
	List(userID string, lastToken int64, limit int) ([]entity.Source, error)
	HaseMore(userID string, lastToken int64, perPage int) (bool, error)

	GetUserSourceByID(ownerID, id string) (*entity.Source, error)
	IsSourceAlreadyCreatedByName(name string) (bool, error)
}

type Service struct {
	repo      SourceRepo
	validator sourcevalidator.Validator
}

func New(repo SourceRepo, validator sourcevalidator.Validator) *Service {
	return &Service{
		repo:      repo,
		validator: validator,
	}
}
