package sourcevalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateIDToDelete(id string) *ValidatorError {
	if err := validation.Validate(id,
		validation.By(v.isSourceExist),
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

		return &ValidatorError{
			Fields: fieldErr,
			Err: richerror.New("sourcevalidator.ValidateIDToDelete").WithMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": id}).WithWrappedError(err),
		}
	}

	return nil
}
