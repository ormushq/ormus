package uservalidator

import (
	"errors"
	"regexp"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/regex"
	"github.com/ormushq/ormus/pkg/richerror"
)

// ValidateLoginRequest is used to validate login request.
func (v Validator) ValidateLoginRequest(req param.LoginRequest) *validator.Error {
	minPasswordLength := 8
	maxPasswordLength := 32

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Match(regexp.MustCompile(regex.Email)).Error(errmsg.ErrEmailIsNotValid), validation.By(v.isUserRegistered)),
		validation.Field(&req.Password, validation.Required, validation.Length(minPasswordLength, maxPasswordLength), validation.By(v.isPasswordValid))); err != nil {

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
			Err: richerror.New("userValidation.ValidateLoginRequest").WhitMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": req}).WhitWarpError(err),
		}
	}

	return nil
}

// isUserRegistered is a helper function in LoginRequest validation process which will return nil if email exists in db and return error otherwise.
func (v Validator) isUserRegistered(value interface{}) error {
	email, ok := value.(string)
	if !ok {
		return richerror.New("validator.isUserRegistered").WhitMessage("wrong type")
	}
	exists, err := v.repo.DoesUserExistsByEmail(email)
	if err != nil {
		return richerror.New("validator.isUserRegistered").WhitWarpError(err).WhitMessage(errmsg.ErrSomeThingWentWrong)
	}

	if !exists {
		return richerror.New("validator.isUserRegistered").WhitMessage(errmsg.ErrAuthUserNotFound)
	}

	return nil
}

// TODO: implement this function with regex
// isPasswordValid is a helper function to validate  password.
func (v Validator) isPasswordValid(value interface{}) error {
	password, ok := value.(string)
	if !ok {
		return richerror.New("validator.isPasswordValid").WhitMessage("wrong type")
	}

	var lower, upper, numeric, special bool

	for _, val := range password {
		switch {
		case unicode.IsNumber(val):
			numeric = true
		case unicode.IsLower(val):
			lower = true
		case unicode.IsUpper(val):
			upper = true
		case unicode.IsSymbol(val) || unicode.IsPunct(val):
			special = true

		}
	}

	if numeric && lower && upper && special {
		return nil
	}

	return richerror.New("validator.isPasswordValid").WhitMessage(errmsg.ErrPasswordIsNotValid)
}
