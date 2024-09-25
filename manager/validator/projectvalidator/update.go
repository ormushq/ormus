package projectvalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

/*
	Logic ValidateUpdateRequest

	1. Name is required
	2. Description is required

*/

func (v Validator) ValidateUpdateRequest(req projectparam.UpdateRequest) *validator.Error {
	const op = "projectvalidator.ValidateUpdateRequest"

	const minNameLength = 3
	const maxNameLength = 0

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minNameLength, maxNameLength)),
		validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.Description, validation.Required),
		validation.Field(&req.ProjectID, validation.Required, validation.By(v.isProjectExist)),
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

// isUserRegistered is a helper function in LoginRequest validation process which will return nil if email exists in db and return error otherwise.
func (v Validator) isProjectExist(value interface{}) error {
	projectID, ok := value.(string)
	if !ok {
		return richerror.New("validator.isProjectExist").WithMessage("wrong type")
	}
	_, err := v.repo.GetWithID(projectID)
	if err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return richerror.New("validator.isProjectExist").WithMessage(errmsg.ErrProjectNotFound)
		}

		return richerror.New("validator.isProjectExist").WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return nil
}
