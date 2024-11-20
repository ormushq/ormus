package sourceservice

import (
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Show(req sourceparam.ShowRequest) (sourceparam.ShowResponse, error) {
	const op = "sourceService.Show"

	vErr := s.validator.ValidateShowRequest(req)
	if vErr != nil {
		return sourceparam.ShowResponse{}, vErr
	}
	source, err := s.repo.GetWithID(req.SourceID)
	if err != nil {
		return sourceparam.ShowResponse{}, richerror.New(op).WithWrappedError(err)
	}

	if source.OwnerID != req.UserID {
		return sourceparam.ShowResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	return sourceparam.ShowResponse{
		Source: source,
	}, nil
}
