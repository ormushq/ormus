package auth_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/service/auth"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"

	"github.com/stretchr/testify/assert"
)

func Test_ParseToken(t *testing.T) {
	testCases := []struct {
		name         string
		bearerToken  string
		expectedErr  error
		expectedUser string // Add more fields if necessary for validation
	}{
		// TODO: these tests may fail due to expiration of the jwt tokens
		{
			name:         "valid token",
			bearerToken:  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhYyIsImV4cCI6MTcwMzE1ODkzNSwidXNlcl9lbWFpbCI6InRlc3RlbWFpbEBleGFtcGxlLmNvbSIsInJvbGUiOiJhZG1pbiJ9.uGyl2lTwhH8CB5UwzYu_2cDrH5zo9_2cCYqivHTY0Cc",
			expectedUser: "testemail@example.com",
		},
		{
			// TODO: this test cases passes but is this correct?
			name:         "valid token without bearer keyword",
			bearerToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhYyIsImV4cCI6MTcwMzE1Nzk2MSwidXNlcl9lbWFpbCI6InRlc3RlbWFpbEBleGFtcGxlLmNvbSJ9.TSvkTtpw69PCjBhqeUQ5t72HHw5GXsyPZrZdETcZYgA",
			expectedUser: "testemail@example.com",
		},
		{
			// the data here is tampered and the jwt module should raise errmsg on this
			name:        "malformed signature",
			bearerToken: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhYyIsImV4cCI6MTcwMzE2MDIxNCwidXNlcl9lbWFpbCI6InRlc3RlbWFpbEBleGFtcGxlLmNvbSIsInJvbGUiOiJhZG1pbiJ9.714GxZ2iXAD87R5Zk27XXgwp7heWYx2190_GAZagFCg",
			expectedErr: fmt.Errorf("token signature is invalid: signature is invalid"), // Define the expected errmsg
		},
	}

	cfg := config.C()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := auth.NewJWT(cfg.Manager.JWTConfig)

			// 2. execution
			claims, err := jwt.ParseToken(tc.bearerToken)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr.Error(), err.(richerror.RichError).Message())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, claims)
				assert.Equal(t, tc.expectedUser, claims.UserEmail)
				// Add more assertions if needed for other claim data
			}
		})
	}
}

func Test_CreateAccessToken(t *testing.T) {
	testCases := []struct {
		name string
		user entity.User
		err  error
	}{
		{
			name: "simple",
			user: entity.User{Email: "testemail@example.com"},
		},
		{
			name: "empty user",
			user: entity.User{},
			err:  richerror.New("Test_CreateAccessToken").WhitMessage(errmsg.ErrJwtEmptyUser),
		},
		{
			name: "nil user",
			err:  richerror.New("Test_CreateAccessToken").WhitMessage(errmsg.ErrJwtEmptyUser),
		},
	}

	cfg := config.C()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := auth.NewJWT(cfg.Manager.JWTConfig)

			// 2. execution
			token, err := jwt.CreateAccessToken(tc.user)
			fmt.Println(token)

			// 3. assertion
			if err != nil {
				assert.Error(t, tc.err, err)
				return
			}
			assert.NotEmpty(t, token)
			assert.NoError(t, err)
		})
	}
}
