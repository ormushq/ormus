package sourceservice

import (
	"github.com/ormushq/ormus/manager/managerparam"
)

func (s Service) UpdateSource(ownerID, sourceID string, req *managerparam.UpdateSourceRequest) (*managerparam.UpdateSourceResponse, error) {
	panic("implement me")
	//source, err := s.repo.GetUserSourceByID(ownerID, sourceID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//source.Name = req.Name
	//source.Description = req.Description
	//source.ProjectID = req.ProjectID
	//source.Status = req.Status
	//
	//response, err := s.repo.UpdateSource(sourceID, source)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return response, nil
}
