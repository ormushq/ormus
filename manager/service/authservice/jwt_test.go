package authservice_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {
	t.Run("test durations", func(t *testing.T) {
		// 1. setup
		cfg := authservice.JwtConfig{
			AccessExpirationTimeInDay:  7,  // 7 * 24 * 60 * 60 * 1000 * 1000 = 604,800,000,000,000
			RefreshExpirationTimeInDay: 28, // 28 * 24 * 60 * 60 * 1000 * 1000 = 2,419,200,000,000,000
		}

		jwt := authservice.NewJWT(cfg)

		// 2. execution
		jwtCfg := jwt.GetConfig()

		// 3. assertion
		assert.Equal(t, time.Duration(604_800_000_000_000), jwtCfg.AccessExpirationTimeInDay)
		assert.Equal(t, time.Duration(2_419_200_000_000_000), jwtCfg.RefreshExpirationTimeInDay)
	})
}

func Test_ParseToken(t *testing.T) {
	testCases := []struct {
		name         string
		bearerToken  string
		expectedErr  error
		expectedUser string // Add more fields if necessary for validation
	}{
		{
			name:         "expired token",
			bearerToken:  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhYyIsImV4cCI6MTcwMzE1ODkzNSwidXNlcl9lbWFpbCI6InRlc3RlbWFpbEBleGFtcGxlLmNvbSIsInJvbGUiOiJhZG1pbiJ9.uGyl2lTwhH8CB5UwzYu_2cDrH5zo9_2cCYqivHTY0Cc",
			expectedUser: "testemail@example.com",
			expectedErr:  fmt.Errorf("token has invalid claims: token is expired"),
		},
		{
			// the data here is tampered and the jwt module should raise error on this
			name:        "malformed signature",
			bearerToken: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhYyIsImV4cCI6MTcwMzE2MDIxNCwidXNlcl9lbWFpbCI6InRlc3RlbWFpbEBleGFtcGxlLmNvbSIsInJvbGUiOiJhZG1pbiJ9.714GxZ2iXAD87R5Zk27XXgwp7heWYx2190_GAZagFCg",
			expectedErr: fmt.Errorf("token signature is invalid: signature is invalid"), // Define the expected error
		},
	}

	cfg := config.C()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := authservice.NewJWT(cfg.Manager.JWTConfig)

			// 2. execution
			claims, err := jwt.ParseToken(tc.bearerToken)

			// 3. assertion
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr.Error(), err.(richerror.RichError).Message())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, claims)
				assert.Equal(t, tc.expectedUser, claims.EnUserEmail)
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
			jwt := authservice.NewJWT(cfg.Manager.JWTConfig)

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
