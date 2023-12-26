package sourceservice

import "github.com/ormushq/ormus/manager/entity"

type SourceRepo interface {
	InsertSource(source *entity.Source) error
	UpdateSource(id string, source *entity.Source) error
	DeleteSource(id string) error
	GetUserSourceById(ownerID, id string) (*entity.Source, error)
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
