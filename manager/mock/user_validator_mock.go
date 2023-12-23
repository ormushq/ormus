package usermock

import (
	"fmt"

	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/param"
)

const ValidatorErr = "validation error"

type MockValidator struct {
	validationErr bool
}

func NewMockValidator(validationErr bool) *MockValidator {
	return &MockValidator{validationErr: validationErr}
}

func (m MockValidator) ValidateLoginRequest(_ param.LoginRequest) *uservalidator.ValidatorError {
	if m.validationErr {
		return &uservalidator.ValidatorError{
			Err: fmt.Errorf(ValidatorErr),
		}
	}

	return nil
}

func (m MockValidator) ValidateRegisterRequest(_ param.RegisterRequest) *uservalidator.ValidatorError {
	if m.validationErr {
		return &uservalidator.ValidatorError{
			Err: fmt.Errorf(ValidatorErr),
		}
	}

	return nil
}
