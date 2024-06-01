package projectservice

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Create() error {
	const op = "projectService.CreateDefaultProject"

	inOutChan, err := s.internalBroker.GetOutputChannel("CreateDefaultProject")
	if err != nil {
		return err
	}
	for msg := range inOutChan {
		_, err := s.repo.Create("Default Project", string(msg.Body))
		if err != nil {
			return richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
		}

		err = msg.Ack()
		if err != nil {
			return richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
		}
		logger.L().Debug("Default project created")
	}

	return nil
}
