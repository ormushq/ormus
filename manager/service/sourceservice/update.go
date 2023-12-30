package sourceservice

import (
	"github.com/ormushq/ormus/manager/param"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s *Service) UpdateSource(id string, req *param.UpdateSourceRequest) (*param.UpdateSourceResponse, error) {
	source, err := s.repo.GetUserSourceByID(req.OwnerID, id)
	if err != nil {
		return nil, richerror.New("UpdateSource").WhitWarpError(err)
	}

	source.Name = req.Name
	source.Description = req.Description
	source.ProjectID = req.ProjectID
	source.Status = req.Status

	response, err := s.repo.UpdateSource(id, source)
	if err != nil {
		return nil, richerror.New("UpdateSource").WhitWarpError(err)
	}

	return response, nil
}
