package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/param"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s *Service) CreateSource(req *param.AddSourceRequest) (*param.AddSourceResponse, error) {
	source := &entity.Source{
		ID:          "",                  // TODO id is ulid ?
		WriteKey:    entity.WriteKey(""), // TODO wait for our write key generator
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     req.OwnerID,
		ProjectID:   req.ProjectID,
	}

	response, err := s.repo.InsertSource(source)
	if err != nil {
		return nil, richerror.New("CreateSource").WhitWarpError(err)
	}

	return response, nil
}
