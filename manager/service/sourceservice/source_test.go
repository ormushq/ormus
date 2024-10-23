package sourceservice_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/mockRepo/sourcemock"
	"github.com/ormushq/ormus/manager/mockRepo/usermock"
	"github.com/ormushq/ormus/manager/service/sourceservice"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSource(t *testing.T) {
	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		req         string
	}{
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("MockRepo.DeleteSource").WithWrappedError(fmt.Errorf(usermock.RepoErr)),
			req:         "source_id",
		},
		{
			name: "ordinary",
			req:  "source_id",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			mockRepo := sourcemock.NewMockRepository(tc.repoErr)
			service := sourceservice.New(mockRepo)

			// 2. execution
			err := service.Delete(tc.req)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)

				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestUpdateSource(t *testing.T) {
	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		sourceID    string
		ownerID     string
		req1        managerparam.UpdateSourceRequest
	}{
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("MockRepo.GetUserSourceById").WithWrappedError(fmt.Errorf(usermock.RepoErr)),
			sourceID:    "source_id",
			ownerID:     "owner_id",
			req1: managerparam.UpdateSourceRequest{
				Name:        "new name",
				Description: "new description",
				ProjectID:   "new project id",
			},
		},
		{
			name:     "ordinary",
			repoErr:  false,
			sourceID: "source_id",
			ownerID:  "owner_id",
			req1: managerparam.UpdateSourceRequest{
				Name:        "new name",
				Description: "new description",
				ProjectID:   "new project id",
			},
		},
		{
			name:        "user not found",
			repoErr:     false,
			expectedErr: richerror.New("MockRepo.GetUserSourceById").WithMessage(errmsg.ErrUserNotFound),
			sourceID:    "invalide source_id",
			ownerID:     "owner_id",
			req1: managerparam.UpdateSourceRequest{
				Name:        "new name",
				Description: "new description",
				ProjectID:   "new project id",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			mockRepo := sourcemock.NewMockRepository(tc.repoErr)
			service := sourceservice.New(mockRepo)

			// 2. execution
			response, err := service.UpdateSource(tc.ownerID, tc.sourceID, &tc.req1)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
				assert.Empty(t, response)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, response)
		})
	}
}

func TestCreateSource(t *testing.T) {
	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		ownerID     string
		req         managerparam.AddSourceRequest
	}{
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("MockRepo.InsertSource").WithWrappedError(fmt.Errorf(usermock.RepoErr)),
			ownerID:     "owner_id",
			req: managerparam.AddSourceRequest{
				Name:        "name",
				Description: "description",
				ProjectID:   "project id",
			},
		},
		{
			name:    "ordinary",
			repoErr: false,
			ownerID: "owner_id",
			req: managerparam.AddSourceRequest{
				Name:        "un existed name",
				Description: "description",
				ProjectID:   "project id",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			mockRepo := sourcemock.NewMockRepository(tc.repoErr)
			service := sourceservice.New(mockRepo)

			// 2. execution
			response, err := service.CreateSource(&tc.req, tc.ownerID)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
				assert.Empty(t, response)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, response)
		})
	}
}
