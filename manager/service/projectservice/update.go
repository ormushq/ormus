package projectservice

import (
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Update(req projectparam.UpdateRequest) (projectparam.UpdateResponse, error) {
	const op = "projectService.Update"

	vErr := s.validator.ValidateUpdateRequest(req)
	if vErr != nil {
		return projectparam.UpdateResponse{}, vErr
	}
	project, err := s.repo.GetWithID(req.ProjectID)
	if err != nil {
		return projectparam.UpdateResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}
	if project.UserID != req.UserID {
		return projectparam.UpdateResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	project.Name = req.Name
	project.Description = req.Description

	project, err = s.repo.Update(project)
	if err != nil {
		return projectparam.UpdateResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return projectparam.UpdateResponse{
		Project: project,
	}, nil
}
