package sourceservice

import (
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/param"
)

func (s *Service) CreateSource(req *param.AddSourceRequest) (*param.AddSourceResponse, error) {

	source := new(entity.Source)
	source.WriteKey = ulid.Make() // TODO wait for our write key generator
	source.Name = req.Name
	source.Description = req.Description
	source.OwnerId = req.OwnerId
	source.ProjectId = req.ProjectId
	source.Status = false // it means source not active
	now := time.Now()
	source.CreateAt = now
	source.UpdateAt = now

	if err := s.repo.InsertSource(source); err != nil {
		return nil, err
	}

	responce := new(param.AddSourceResponse)
	responce.WriteKey = source.WriteKey
	responce.Name = source.Name
	responce.Description = source.Description
	responce.OwnerId = source.OwnerId
	responce.ProjectId = source.ProjectId
	responce.Status = source.Status
	responce.CreateAt = source.CreateAt
	responce.UpdateAt = source.UpdateAt

	return responce, nil
}

func (s *Service) UpdateSource(id string, req *param.UpdateSourceRequest) (*param.UpdateSourceResponse, error) {

	source, err := s.repo.GetUserSources(req.OwnerId, id)
	if err != nil {
		return nil, err
	}

	source.Name = req.Name
	source.Description = req.Description
	source.ProjectId = req.ProjectId
	source.Status = req.Status
	now := time.Now()
	source.UpdateAt = now

	if err := s.repo.UpdateSource(id, source); err != nil {
		return nil, err
	}

	responce := new(param.UpdateSourceResponse)
	responce.WriteKey = source.WriteKey
	responce.Name = source.Name
	responce.Description = source.Description
	responce.OwnerId = source.OwnerId
	responce.ProjectId = source.ProjectId
	responce.Status = source.Status
	responce.CreateAt = source.CreateAt
	responce.UpdateAt = source.UpdateAt

	return responce, nil
}

func (s *Service) DeleteSource(id string) error {

	if err := s.repo.DeleteSource(id); err != nil {
		return err
	}

	return nil
}
