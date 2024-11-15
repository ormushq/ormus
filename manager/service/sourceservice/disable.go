package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Disable(req sourceparam.DisableRequest) (sourceparam.DisableResponse, error) {
	const op = "sourceService.Disable"

	vErr := s.validator.ValidateDisableRequest(req)
	if vErr != nil {
		return sourceparam.DisableResponse{}, vErr
	}
	source, err := s.repo.GetWithID(req.SourceID)
	if err != nil {
		return sourceparam.DisableResponse{}, richerror.New(op).WithWrappedError(err)
	}

	if source.OwnerID != req.UserID {
		return sourceparam.DisableResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	source.Status = entity.SourceStatusNotActive

	_, err = s.repo.Update(source)
	if err != nil {
		return sourceparam.DisableResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return sourceparam.DisableResponse{
		Message: "source Disabled",
	}, nil
}
