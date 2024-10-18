package projectservice

import (
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Get(req projectparam.GetRequest) (projectparam.GetResponse, error) {
	const op = "projectService.Get"

	vErr := s.validator.ValidateGetRequest(req)
	if vErr != nil {
		return projectparam.GetResponse{}, vErr
	}

	project, err := s.repo.GetWithID(req.ProjectID)
	if err != nil {
		return projectparam.GetResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}
	if project.UserID != req.UserID {
		return projectparam.GetResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	return projectparam.GetResponse{
		Project: project,
	}, nil
}
