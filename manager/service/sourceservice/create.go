package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/param"
	writekey "github.com/ormushq/ormus/pkg/write_key"
)

func (s *Service) CreateSource(req *param.AddSourceRequest) (*param.AddSourceResponse, error) {
	w, err := writekey.GenerateNewWriteKey()
	if err != nil {
		return nil, err
	}

	source := &entity.Source{
		ID:          "", // TODO  uuid ulid ?
		WriteKey:    entity.WriteKey(w),
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     req.OwnerID,
		ProjectID:   req.ProjectID,
	}

	response, err := s.repo.InsertSource(source)
	if err != nil {
		return nil, err
	}

	return response, nil
}
