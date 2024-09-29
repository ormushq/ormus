package sourcevalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateIDToDelete(id string) *ValidatorError {
	if err := validation.Validate(id,
		validation.By(v.isSourceAlreadyCreated),
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

func (v Validator) isSourceAlreadyCreated(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	exist, err := v.repo.IsSourceAlreadyCreatedByName(s)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("this name is already usesd")
	}

	return nil
}
