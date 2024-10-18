package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/manager/validator/sourcevalidator"
)

type SourceRepo interface {
	Create(source entity.Source) (entity.Source, error)
	GetWithID(id string) (entity.Source, error)
	Update(source entity.Source) (entity.Source, error)
	Delete(source entity.Source) error
	List(ownerID string, lastToken int64, limit int) ([]entity.Source, error)
	HaseMore(ownerID string, lastToken int64, perPage int) (bool, error)
}

type ProjectSvc interface {
	Get(projectparam.GetRequest) (projectparam.GetResponse, error)
}
type Service struct {
	repo       SourceRepo
	validator  sourcevalidator.Validator
	projectSvc ProjectSvc
}

func New(repo SourceRepo, validator sourcevalidator.Validator, projectSvc ProjectSvc) Service {
	return Service{
		repo:       repo,
		validator:  validator,
		projectSvc: projectSvc,
	}
}
