package sourcevalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateDeleteRequest(req sourceparam.DeleteRequest) *validator.Error {
	const op = "sourcevalidator.ValidateDeleteRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.SourceID, validation.Required, validation.By(v.isSourceExist)),
	); err != nil {
		fieldErr := make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)

		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErr[key] = value.Error()
				}
			}
		}

		return &validator.Error{
			Fields: fieldErr,
			Err: richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": req}).WithWrappedError(err),
		}
	}

	return nil
}
