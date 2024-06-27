package uservalidator

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/regex"
	"github.com/ormushq/ormus/pkg/richerror"
)

// TODO: should we change the name to just `Register`?

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) *validator.Error {
	minNameLength := 3
	maxNameLength := 50

	minPasswordLength := 8
	maxPasswordLength := 32

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minNameLength, maxNameLength)),
		validation.Field(&req.Email, validation.Required, validation.Match(regexp.MustCompile(regex.Email)).Error(errmsg.ErrEmailIsNotValid), validation.By(v.isEmailAlreadyRegistered)),
		validation.Field(&req.Password, validation.Required, validation.Length(minPasswordLength, maxPasswordLength), validation.By(v.isPasswordValid)),
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
			Err: richerror.New("validation.register").WithMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": req}).WithWrappedError(err),
		}
	}

	return nil
}

// isEmailAlreadyRegistered is a function which will check if user is already registered in the register process and will return error if user was existing and registered and nil if was not registered
// TODO: isEmailAlreadyRegistered and isUserRegistered are similar in most parts but at the end of function we have different logics, so I decided to repeat the code and sacrifice redundancy for readability.
//
//	func (v Validator) isEmailAlreadyRegistered(value interface{}) error {
//		err := v.isUserRegistered(value)
//		if err == nil {
//			return richerror.New("validator.isEmailAlreadyRegistered").WithMessage(errmsg.ErrAuthUserExisting)
//		}
//
//		return nil
//	}
//
// you can implement it like upper code if you think it is better.
func (v Validator) isEmailAlreadyRegistered(value interface{}) error {
	email, ok := value.(string)
	if !ok {
		return richerror.New("validator.isEmailAlreadyRegistered").WithMessage("wrong type")
	}
	exists, err := v.repo.DoesUserExistsByEmail(email)
	if err != nil {
		return richerror.New("validator.isEmailAlreadyRegistered").WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	if exists {
		return richerror.New("validator.isEmailAlreadyRegistered").WithMessage(errmsg.ErrAuthUserExisting)
	}

	return nil
}
