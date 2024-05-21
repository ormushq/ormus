package projectvalidator_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/validator/projectvalidator"
	"github.com/ormushq/ormus/param"
	"github.com/stretchr/testify/assert"
)

func TestValidator_ValidateCreateRequest(t *testing.T) {
	testCases := []struct {
		name      string
		params    param.CreateProjectRequest
		svcErr    bool
		validUser bool
		error     error
	}{
		{
			name: "happy path",
			params: param.CreateProjectRequest{
				Name:   "correct name",
				UserID: "0000000000",
			},
			validUser: true,
		},
		{
			name:  "name is required",
			error: fmt.Errorf("Name: cannot be blank\n"),
			params: param.CreateProjectRequest{
				UserID: "0000000000",
			},
			validUser: true,
		},
		{
			name: "name shorter than 3",
			params: param.CreateProjectRequest{
				Name:   "co",
				UserID: "0000000000",
			},
			error:     fmt.Errorf("Name: the length must be no less than 3\n"),
			validUser: true,
		},
		{
			name:  "user_id is required",
			error: fmt.Errorf("UserID: cannot be blank\n"),
			params: param.CreateProjectRequest{
				Name: "correct name",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := NewStubUserExistenceChecker(tc.svcErr, tc.validUser)
			vld := projectvalidator.New(repo)

			// 2. execution
			res := vld.ValidateCreateRequest(tc.params)

			// 3. assertion
			if tc.error == nil {
				assert.Nil(t, res)
				return
			}
			assert.Equal(t, tc.error.Error(), res.Error())
		})
	}
}

const ServiceErr = "service error"

type StubUserExistenceChecker struct {
	hasErr    bool
	validUser bool
}

func NewStubUserExistenceChecker(hasErr bool, validUser bool) StubUserExistenceChecker {
	return StubUserExistenceChecker{
		hasErr:    hasErr,
		validUser: validUser,
	}
}

// IsUserIDValid is a stub implementation of the IsUserIDValid method.
func (s StubUserExistenceChecker) IsUserIDValid(userID string) (bool, error) {
	if s.hasErr {
		return false, fmt.Errorf(ServiceErr)
	}

	if s.validUser {
		return true, nil
	}

	return false, nil
}
