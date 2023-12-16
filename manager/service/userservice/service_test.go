package service_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/service/auth"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/stretchr/testify/assert"
)

func TestService_Register(t *testing.T) {
	// TODO: if password is longer than 72 bycrypt will fail

	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		req         param.RegisterRequest
	}{
		{
			name:        "user exists",
			expectedErr: richerror.New("register").WhitMessage(errmsg.ErrAuthUserExisting),
			req: param.RegisterRequest{
				Email:    "test@example.com",
				Password: "123",
			},
		},
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("register.repo").WhitWarpError(fmt.Errorf(errRepo)),
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
			repo := NewMockRepository(tc.repoErr)
			svc := service.NewService(jwt, repo)

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
	testCases := []struct {
		name        string
		repoErr     bool
		expectedErr error
		req         param.LoginRequest
	}{
		{
			name:        "user not available",
			expectedErr: richerror.New("Login").WhitMessage(errmsg.ErrWrongCredentials),
			req: param.LoginRequest{
				Email:    "not@existing.com",
				Password: "123",
			},
		},
		{
			name:        "wrong password",
			expectedErr: richerror.New("Login").WhitMessage(errmsg.ErrWrongCredentials),
			req: param.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
		},
		{
			name:        "repo fails",
			repoErr:     true,
			expectedErr: richerror.New("Login").WhitWarpError(fmt.Errorf(errRepo)),
			req: param.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
		},
		{
			name: "ordinary",
			req: param.LoginRequest{
				Email:    "test@example.com",
				Password: "very_strong_password",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := MockJwtEngine{}
			repo := NewMockRepository(tc.repoErr)
			svc := service.NewService(jwt, repo)

			// 2. execution
			user, err := svc.Login(tc.req)
			if err != nil {
				return
			}

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
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

const errRepo = "repository error"

type MockRepository struct {
	users  []entity.User
	hasErr bool
}

func NewMockRepository(hasErr bool) *MockRepository {
	users := []entity.User{
		{
			Email:    "test@example.com",
			Password: "$2a$10$pMV1Q1b9jgUQdgWnq4GVOuenS2X.HPns0oMRYmoLdLR1nJL/oONzS", // very_strong_password
		},
	}

	return &MockRepository{
		users:  users,
		hasErr: hasErr,
	}
}

func (m MockRepository) Register(u entity.User) (*entity.User, error) {
	if m.hasErr {
		return nil, fmt.Errorf(errRepo)
	}

	u.ID = "new_id"
	m.users = append(m.users, u)

	return &u, nil
}

func (m MockRepository) GetUserByEmail(email string) (*entity.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, richerror.New("MockRepo.GetUserByEmail").WhitMessage(errmsg.ErrAuthUserNotFound)
}

func (m MockRepository) DoesUserExistsByEmail(email string) (bool, error) {
	for _, user := range m.users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}
