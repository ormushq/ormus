package sourcevalidator

import (
	"errors"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateUpdateSourceForm(req sourceparam.UpdateRequest) *validator.Error {
	const op = "sourcevalidator.ValidateUpdateRequest"

	minNameLength := 5
	maxNameLength := 30

	minDescriptionLength := 5
	maxDescriptionLength := 100

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minNameLength, maxNameLength)),
		validation.Field(&req.Description, validation.Required, validation.Length(minDescriptionLength, maxDescriptionLength)),
		validation.Field(&req.UserID, validation.Required, validation.By(v.validateULID)),
		validation.Field(&req.Status, validation.Required, validation.By(v.validateStatus)),
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

func (v Validator) validateStatus(value interface{}) error {
	s, ok := value.(entity.Status)
	if !ok {
		return errors.New("error while reflection interface")
	}

	switch s {
	case "active":
		return nil
	case "not active":
		return nil
	}

	return errors.New("invalide status of source")
}
