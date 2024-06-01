package projectvalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

/*
	Logic ValidateCreateRequest

	1. Name is required
	2. Email is required
	3. Email have to be valid
	4. Email regex have to match

*/

func (v Validator) ValidateCreateRequest(req param.CreateProjectRequest) *validator.Error {
	const op = "projectvalidator.ValidateCreateRequest"

	const minNameLength = 3
	const maxNameLength = 0

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minNameLength, maxNameLength)),
		validation.Field(&req.UserID, validation.Required),
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
