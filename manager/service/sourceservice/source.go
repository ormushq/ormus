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
	source.OwnerID = req.OwnerID
	source.ProjectID = req.ProjectID
	source.Status = false // it means source not active
	now := time.Now()
	source.CreateAt = now
	source.UpdateAt = now

	if err := s.repo.InsertSource(source); err != nil {
		return nil, err
	}

	response := new(param.AddSourceResponse)
	response.WriteKey = source.WriteKey
	response.Name = source.Name
	response.Description = source.Description
	response.OwnerID = source.OwnerID
	response.ProjectID = source.ProjectID
	response.Status = source.Status
	response.CreateAt = source.CreateAt
	response.UpdateAt = source.UpdateAt

	return response, nil
}

func (s *Service) UpdateSource(id string, req *param.UpdateSourceRequest) (*param.UpdateSourceResponse, error) {
	source, err := s.repo.GetUserSourceById(req.OwnerID, id)
	if err != nil {
		return nil, err
	}

	source.Name = req.Name
	source.Description = req.Description
	source.ProjectID = req.ProjectID
	source.Status = req.Status
	now := time.Now()
	source.UpdateAt = now

	if err := s.repo.UpdateSource(id, source); err != nil {
		return nil, err
	}

	response := new(param.UpdateSourceResponse)
	response.WriteKey = source.WriteKey
	response.Name = source.Name
	response.Description = source.Description
	response.OwnerID = source.OwnerID
	response.ProjectID = source.ProjectID
	response.Status = source.Status
	response.CreateAt = source.CreateAt
	response.UpdateAt = source.UpdateAt

	return response, nil
}

func (s *Service) DeleteSource(id string) error {
	if err := s.repo.DeleteSource(id); err != nil {
		return err
	}

	return nil
}
