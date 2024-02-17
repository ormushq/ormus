package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/param"
)

type SourceRepo interface {
	InsertSource(source *entity.Source) (*param.AddSourceResponse, error)
	UpdateSource(id string, source *entity.Source) (*param.UpdateSourceResponse, error)
	DeleteSource(id, userID string) error
	GetUserSourceByID(ownerID, id string) (*entity.Source, error)
	IsSourceAlreadyCreatedByName(name string) (bool, error)
}

type Service struct {
	repo SourceRepo
}

func New(repo SourceRepo) *Service {
	return &Service{
		repo: repo,
	}
}
