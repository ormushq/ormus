package userservice_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/mockRepo/projectstub"
	"github.com/ormushq/ormus/manager/mockRepo/usermock"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/manager/workers"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/simple"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/stretchr/testify/assert"
)

func TestService_Register(t *testing.T) {
	// TODO: if password is longer than 72 bycrypt will fail
	cfg := config.C().Manager
	done := make(chan bool)
	wg := sync.WaitGroup{}
	internalBroker := simple.New(done, &wg)
	internalBroker.NewChannel("CreateDefaultProject", channel.BothMode,
		cfg.InternalBrokerConfig.ChannelSize, cfg.InternalBrokerConfig.NumberInstant, cfg.InternalBrokerConfig.MaxRetryPolicy)
	RepoPr := projectstub.New(false)
	ProjectSvc := projectservice.New(&RepoPr, internalBroker)
	workers.New(ProjectSvc, internalBroker).Run(done, &wg)
	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		req         param.RegisterRequest
	}{
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("register.repo").WithWrappedError(fmt.Errorf(usermock.RepoErr)),
			req: param.RegisterRequest{
				Email:    "new@example.com",
				Password: "very_safe_password",
			},
		},
		{
			name: "ordinary",
			req: param.RegisterRequest{
				Email:    "new@user.com",
				Password: "very_safe_password",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := MockJwtEngine{}
			repo := usermock.NewMockRepository(tc.repoErr)
			validateUserSvc := uservalidator.New(repo)
			svc := userservice.New(jwt, repo, internalBroker, validateUserSvc)

			// 2. execution
			user, err := svc.Register(tc.req)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
				assert.Empty(t, user)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, user)
		})
	}
}

func TestService_Login(t *testing.T) {
	defaultUser := usermock.DefaultUser()

	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		req         param.LoginRequest
	}{
		{
			name: "ordinary",
			req: param.LoginRequest{
				Email:    defaultUser.Email,
				Password: defaultUser.Password,
			},
		},
		{
			name:        "user not available",
			expectedErr: richerror.New("Login").WithWrappedError(fmt.Errorf(usermock.RepoErr)),
			req: param.LoginRequest{
				Email:    "not@existing.com",
				Password: "123",
			},
		},
		{
			name:        "wrong password",
			expectedErr: richerror.New("Login").WithMessage(errmsg.ErrWrongCredentials),
			req: param.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
		},
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("Login").WithWrappedError(fmt.Errorf(usermock.RepoErr)),
			req: param.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := MockJwtEngine{}

			repo := usermock.NewMockRepository(tc.repoErr)
			validateUserSvc := uservalidator.New(repo)
			svc := userservice.New(jwt, repo, nil, validateUserSvc)

			// 2. execution
			user, err := svc.Login(tc.req)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
				assert.Empty(t, user)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, user)
			assert.NotEmpty(t, user.User.Email)
		})
	}
}

type MockJwtEngine struct{}

func (m MockJwtEngine) CreateAccessToken(user entity.User) (string, error) {
	return "very_secure_token", nil
}

func (m MockJwtEngine) CreateRefreshToken(user entity.User) (string, error) {
	return "very_secure_token", nil
}
