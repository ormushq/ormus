package sourceservice

import (
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) List(req sourceparam.ListRequest) (sourceparam.ListResponse, error) {
	const op = "projectService.List"

	sources, err := s.repo.List(req.UserID, req.LastTokenID, req.PerPage)
	if err != nil {
		return sourceparam.ListResponse{}, richerror.New(op).WithWrappedError(err)
	}

	haseMore, err := s.repo.HaseMore(req.UserID, req.LastTokenID, req.PerPage)
	if err != nil {
		return sourceparam.ListResponse{}, richerror.New(op).WithWrappedError(err)
	}

	return sourceparam.ListResponse{
		Sources:     sources,
		HasMore:     haseMore,
		LastTokenID: req.LastTokenID,
		PerPage:     req.PerPage,
	}, nil
}
