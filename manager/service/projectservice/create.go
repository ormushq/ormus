package projectservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Create(req projectparam.CreateRequest) (projectparam.CreateResponse, error) {
	const op = "projectService.Create"

	vErr := s.validator.ValidateCreateRequest(req)
	if vErr != nil {
		return projectparam.CreateResponse{}, vErr
	}

	project := entity.Project{
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserID,
	}
	project, err := s.repo.Create(project)
	if err != nil {
		return projectparam.CreateResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return projectparam.CreateResponse{
		Project: project,
	}, nil
}
