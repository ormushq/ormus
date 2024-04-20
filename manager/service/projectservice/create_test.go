package projectservice_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/mock/projectstub"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/stretchr/testify/assert"
)

func TestService_Create(t *testing.T) {
	testCases := []struct {
		name    string
		repoErr bool
		req     param.CreateProjectRequest
		err     error
	}{
		{
			name: "happy path",
			req: param.CreateProjectRequest{
				Name:   "new project",
				UserID: "0000000000",
			},
		},
		{
			name:    "repo error",
			repoErr: true,
			err:     fmt.Errorf(errmsg.ErrSomeThingWentWrong),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			repo := projectstub.New(tc.repoErr)
			svc := projectservice.New(&repo)

			// 2. execution
			newProject, err := svc.Create(tc.req)

			// 3. assertion
			if tc.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tc.err.Error(), err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, newProject.Name)
			assert.NotEmpty(t, newProject.ID)
		})
	}
}
