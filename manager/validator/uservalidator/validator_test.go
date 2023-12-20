package uservalidator_test

import (
	"fmt"
	"github.com/ormushq/ormus/manager/mock"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_ValidateLoginRequest(t *testing.T) {
	testCases := []struct {
		name    string
		params  param.LoginRequest
		repoErr bool
		error   error
	}{
		{
			name: "ordinary login",
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "HeavY!234",
			},
		},
		{
			name:  "email cannot be empty",
			error: fmt.Errorf(fmt.Sprintf("email: cannot be blank\n")),
			params: param.LoginRequest{
				Email:    "",
				Password: "HeavY!234",
			},
		},
		{
			name:  "email regex fail",
			error: fmt.Errorf(fmt.Sprintf("email: %s\n", errmsg.ErrEmailIsNotValid)),
			params: param.LoginRequest{
				Email:    "wrongemail.com",
				Password: "HeavY!234",
			},
		},
		{
			name:  "email regex fail 2",
			error: fmt.Errorf(fmt.Sprintf("email: %s\n", errmsg.ErrEmailIsNotValid)),
			params: param.LoginRequest{
				Email:    "wrongemail@salam",
				Password: "HeavY!234",
			},
		},
		{
			name:  "user with email does not exists",
			error: fmt.Errorf(fmt.Sprintf("email: %s\n", errmsg.ErrAuthUserNotFound)),
			params: param.LoginRequest{
				Email:    "test1@example.com",
				Password: "HeavY!234",
			},
		},
		{
			name:  "password should not be empty",
			error: fmt.Errorf(fmt.Sprintf("password: cannot be blank\n")),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "",
			},
		},
		{
			name:  "password length have to be longer than 8",
			error: fmt.Errorf(fmt.Sprintf("password: %s\n", errmsg.ErrPasswordIsTooShort)),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "Sa!1",
			},
		},
		{
			name:  "password length have to be shorter than 32",
			error: fmt.Errorf(fmt.Sprintf("password: %s\n", errmsg.ErrPasswordIsTooLong)),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "Sa!1DJKFDJFKJFakdjfkdsjfjkSDkfjksjdfksKSDFKJSFasdjflajsdflkjsdfkLKSJdf",
			},
		},
		{
			name:  "password have to include number",
			error: fmt.Errorf(fmt.Sprintf("password: %s\n", errmsg.ErrPasswordIsNotValid)),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "Sa!Sa!Sa!Sa!",
			},
		},
		{
			name:  "password have to include lower case letter",
			error: fmt.Errorf(fmt.Sprintf("password: %s\n", errmsg.ErrPasswordIsNotValid)),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "SA!SA!SA!SA!123",
			},
		},
		{
			name:  "password have to include upper case letter",
			error: fmt.Errorf(fmt.Sprintf("password: %s\n", errmsg.ErrPasswordIsNotValid)),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "sa!sa!sa!sa!123",
			},
		},
		{
			name:  "password have to include symbols",
			error: fmt.Errorf(fmt.Sprintf("password: %s\n", errmsg.ErrPasswordIsNotValid)),
			params: param.LoginRequest{
				Email:    "test@example.com",
				Password: "Sa1Sa1Sa1Sa1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := userrepomock_test.NewMockRepository(tc.repoErr)
			vld := uservalidator.New(repo)

			// 2. execution
			res := vld.ValidateLoginRequest(tc.params)

			// 3. assertion
			if tc.error == nil {
				assert.Nil(t, res)
				return
			}
			assert.Equal(t, tc.error.Error(), res.Error())

		})
	}

}
