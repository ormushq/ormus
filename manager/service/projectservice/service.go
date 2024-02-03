package projectservice

import "github.com/ormushq/ormus/manager/entity"

type Repository interface {
	Create(name, email string) (entity.Project, error)
}

type Service struct {
	repo Repository
}

func New(repository Repository) *Service {
	return &Service{repo: repository}
}
