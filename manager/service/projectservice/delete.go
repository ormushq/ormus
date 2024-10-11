package projectservice

import (
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Delete(req projectparam.DeleteRequest) (projectparam.DeleteResponse, error) {
	const op = "projectService.Delete"

	vErr := s.validator.ValidateDeleteRequest(req)
	if vErr != nil {
		return projectparam.DeleteResponse{}, vErr
	}

	project, err := s.repo.GetWithID(req.ProjectID)
	if err != nil {
		return projectparam.DeleteResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}
	if project.UserID != req.UserID {
		return projectparam.DeleteResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	err = s.repo.Delete(project)
	if err != nil {
		return projectparam.DeleteResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return projectparam.DeleteResponse{
		Message: "project deleted successfully",
	}, nil
}
