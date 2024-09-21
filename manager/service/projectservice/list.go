package projectservice

import (
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) List(req projectparam.ListRequest) (projectparam.ListResponse, error) {
	const op = "projectService.List"

	projects, err := s.repo.List(req.UserID, req.LastTokenID, req.PerPage)
	if err != nil {
		return projectparam.ListResponse{}, richerror.New(op).WithWrappedError(err)
	}

	haseMore, err := s.repo.HaseMore(req.UserID, req.LastTokenID, req.PerPage)
	if err != nil {
		return projectparam.ListResponse{}, richerror.New(op).WithWrappedError(err)
	}

	return projectparam.ListResponse{
		Projects:    projects,
		HasMore:     haseMore,
		LastTokenID: req.LastTokenID,
		PerPage:     req.PerPage,
	}, nil
}
