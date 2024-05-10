package projectservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/mock/projectstub"
)

type Repository interface {
	Create(name, email string) (entity.Project, error)
}

type Service struct {
	repo Repository
}

func New(repository projectstub.MockProject) *Service {
	return &Service{repo: &repository}
}
