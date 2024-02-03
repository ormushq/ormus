package projectvalidator

import (
	"errors"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/regex"
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
		validation.Field(&req.UserEmail, validation.Required, validation.Match(regexp.MustCompile(regex.Email)).Error(errmsg.ErrEmailIsNotValid), validation.By(v.isUserEmailValid)),
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
			Err: richerror.New(op).WhitMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": req}).WhitWarpError(err),
		}
	}

	return nil
}

func (v Validator) isUserEmailValid(value interface{}) error {
	userID, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not a valid string")
	}

	doesExists, err := v.userExistenceChecker.IsUserIDValid(userID)
	if err != nil {
		// return service error
		return err
	}

	if doesExists {
		return nil
	}

	return fmt.Errorf(errmsg.ErrEmailIsNotValid)
}
