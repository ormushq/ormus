package sourceservice_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/mock/sourcemock"
	"github.com/ormushq/ormus/manager/param"
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
			expectedErr: richerror.New("MockRepo.DeleteSource").WhitWarpError(fmt.Errorf(sourcemock.RepoErr)),
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
			err := service.DeleteSource(tc.req, "userID")

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
		req1        param.UpdateSourceRequest
	}{
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("MockRepo.GetUserSourceById").WhitWarpError(fmt.Errorf(sourcemock.RepoErr)),
			sourceID:    "source_id",
			ownerID:     "owner_id",
			req1: param.UpdateSourceRequest{
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
			req1: param.UpdateSourceRequest{
				Name:        "new name",
				Description: "new description",
				ProjectID:   "new project id",
			},
		},
		{
			name:        "user not found",
			repoErr:     false,
			expectedErr: richerror.New("MockRepo.GetUserSourceById").WhitMessage(errmsg.ErrUserNotFound),
			sourceID:    "invalide source_id",
			ownerID:     "owner_id",
			req1: param.UpdateSourceRequest{
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
		req         param.AddSourceRequest
	}{
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("MockRepo.InsertSource").WhitWarpError(fmt.Errorf(sourcemock.RepoErr)),
			ownerID:     "owner_id",
			req: param.AddSourceRequest{
				Name:        "name",
				Description: "description",
				ProjectID:   "project id",
			},
		},
		{
			name:    "ordinary",
			repoErr: false,
			ownerID: "owner_id",
			req: param.AddSourceRequest{
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
