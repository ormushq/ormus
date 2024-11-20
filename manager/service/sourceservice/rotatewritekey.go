package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	writekey "github.com/ormushq/ormus/pkg/write_key"
)

func (s Service) RotateWriteKey(req sourceparam.RotateWriteKeyRequest) (sourceparam.RotateWriteKeyResponse, error) {
	const op = "sourceService.RotateWriteKey"

	vErr := s.validator.ValidateRotateWriteKeyRequest(req)
	if vErr != nil {
		return sourceparam.RotateWriteKeyResponse{}, vErr
	}
	source, err := s.repo.GetWithID(req.SourceID)
	if err != nil {
		return sourceparam.RotateWriteKeyResponse{}, richerror.New(op).WithWrappedError(err)
	}

	if source.OwnerID != req.UserID {
		return sourceparam.RotateWriteKeyResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}

	w, err := writekey.GenerateNewWriteKey()
	if err != nil {
		return sourceparam.RotateWriteKeyResponse{}, err
	}
	source.WriteKey = entity.WriteKey(w)
	source, err = s.repo.Update(source)
	if err != nil {
		return sourceparam.RotateWriteKeyResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return sourceparam.RotateWriteKeyResponse{
		Source: source,
	}, nil
}
