package eventvalidator

import (
	"context"
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateWriteKey(ctx context.Context, writeKey string) (bool, error) {

	isValid, err := v.repo.IsWriteKeyValid(ctx, writeKey)
	if err != nil {

		return false, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
	}
	if !isValid {

		Result, err := v.managerValidation.IsWriteKeyValid(ctx, &source.ValidateWriteKeyReq{
			WriteKey: writeKey,
		})
		if err != nil {
			return false, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
		}

		// if the write key is valid, save it into cache
		if Result.IsValid {
			err := v.repo.CreateNewWriteKey(ctx, &source.NewSourceEvent{WriteKey: writeKey}, v.config.WriteKeyRedisExpiration)
			if err != nil {
				return false, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
			}
			return true, nil
		}
	} else {
		// if the write key is valid
		return true, nil
	}

	return false, nil
}
