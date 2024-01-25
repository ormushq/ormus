package projectservice

import (
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Create(req param.CreateProjectRequest) (*param.CreateProjectResponse, error) {
	const op = "projectservice.Create"

	newProject, err := s.repo.Create(req.Name, req.UserEmail)
	if err != nil {
		return nil, richerror.New(op).WhitWarpError(err).WhitMessage(errmsg.ErrSomeThingWentWrong)
	}

	return &param.CreateProjectResponse{
		Name: newProject.Name,
	}, nil
}
