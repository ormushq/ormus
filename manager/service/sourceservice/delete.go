package sourceservice

import (
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Delete(req sourceparam.DeleteRequest) (sourceparam.DeleteResponse, error) {
	const op = "sourceservice.Update"

	vErr := s.validator.ValidateDeleteRequest(req)
	if vErr != nil {
		return sourceparam.DeleteResponse{}, vErr
	}

	source, err := s.repo.GetWithID(req.SourceID)
	if err != nil {
		return sourceparam.DeleteResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}
	if source.OwnerID != req.UserID {
		return sourceparam.DeleteResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	err = s.repo.Delete(source)
	if err != nil {
		return sourceparam.DeleteResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return sourceparam.DeleteResponse{
		Message: "source deleted successfully",
	}, nil
}
