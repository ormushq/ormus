package projectvalidator_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/manager/validator/projectvalidator"
	"github.com/stretchr/testify/assert"
)

func TestValidator_ValidateCreateRequest(t *testing.T) {
	testCases := []struct {
		name    string
		params  projectparam.CreateRequest
		svcErr  bool
		project entity.Project
		error   error
	}{
		{
			name: "happy path",
			params: projectparam.CreateRequest{
				Name:        "correct name",
				UserID:      "test-user-id",
				Description: "Description",
			},
			project: entity.Project{},
		},
		{
			name:  "name is required",
			error: fmt.Errorf("name: cannot be blank\n"),
			params: projectparam.CreateRequest{
				UserID:      "0000000000",
				Description: "Description",
			},
			project: entity.Project{},
		},
		{
			name:  "description is required",
			error: fmt.Errorf("description: cannot be blank\n"),
			params: projectparam.CreateRequest{
				UserID: "0000000000",
				Name:   "correct name",
			},
			project: entity.Project{},
		},

		{
			name: "name shorter than 3",
			params: projectparam.CreateRequest{
				Name:        "co",
				UserID:      "0000000000",
				Description: "Description",
			},
			error:   fmt.Errorf("name: the length must be no less than 3\n"),
			project: entity.Project{},
		},
		{
			name:  "user_id is required",
			error: fmt.Errorf("UserID: cannot be blank\n"),
			params: projectparam.CreateRequest{
				Name:        "correct name",
				Description: "description",
			},
			project: entity.Project{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := NewStubProjectRepository(tc.svcErr, tc.project)
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

type StubProjectRepository struct {
	hasErr  bool
	project entity.Project
}

func NewStubProjectRepository(hasErr bool, project entity.Project) StubProjectRepository {
	return StubProjectRepository{
		hasErr:  hasErr,
		project: project,
	}
}

// IsUserIDValid is a stub implementation of the IsUserIDValid method.
func (s StubProjectRepository) GetWithID(id string) (entity.Project, error) {
	if s.hasErr {
		return s.project, fmt.Errorf(ServiceErr)
	}

	return s.project, nil
}
