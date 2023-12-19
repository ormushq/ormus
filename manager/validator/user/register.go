package uservalidator

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) (map[string]string, error) {
	maxLength := 50
	minLength := 3

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minLength, maxLength)),
		validation.Field(&req.Email, validation.Required, validation.Match(regexp.MustCompile(emailRegex)), validation.By(v.doesUserExistsByEmail)),
		validation.Field(&req.Password, validation.Required, validation.By(v.isPasswordValid)),
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

		return fieldErr, richerror.New("validation.register").WhitMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
			WhitMeta(map[string]interface{}{"request:": req}).WhitWarpError(err)
	}

	return nil, nil
}

// doesUserExistsByEmail it's a helper function checks the user is exists
// this function used for registration users.
func (v Validator) doesUserExistsByEmail(value interface{}) error {
	// fetch user to check if exists before user creation
	email, ok := value.(string)
	if !ok {
		return richerror.New("validator.doesUserExist").WhitMessage("wrong type")
	}

	existing, err := v.repo.DoesUserExistsByEmail(email)
	if err != nil {
		return richerror.New("validation.doesUserExistsByEmail").WhitWarpError(err)
	}
	if existing {
		return richerror.New("validation.doesUserExistsByEmail").WhitMessage(errmsg.ErrAuthUserExisting)
	}

	return nil
}
