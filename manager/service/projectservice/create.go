package projectservice

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Create(msg []byte) error {
	const op = "projectService.CreateDefaultProject"
	_, err := s.repo.Create("Default Project", string(msg))
	if err != nil {
		return richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}
	logger.L().Debug("Default project created")

	return nil
}
