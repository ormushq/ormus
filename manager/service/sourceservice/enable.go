package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Enable(req sourceparam.EnableRequest) (sourceparam.EnableResponse, error) {
	const op = "sourceService.Enable"

	vErr := s.validator.ValidateEnableRequest(req)
	if vErr != nil {
		return sourceparam.EnableResponse{}, vErr
	}
	source, err := s.repo.GetWithID(req.SourceID)
	if err != nil {
		return sourceparam.EnableResponse{}, richerror.New(op).WithWrappedError(err)
	}

	if source.OwnerID != req.UserID {
		return sourceparam.EnableResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	source.Status = entity.SourceStatusActive

	_, err = s.repo.Update(source)
	if err != nil {
		return sourceparam.EnableResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return sourceparam.EnableResponse{
		Message: "source enabled",
	}, nil
}
