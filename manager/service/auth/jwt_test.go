package auth_test

import (
	"fmt"
	"github.com/ormushq/ormus/config"
	"testing"

	"github.com/ormushq/ormus/manager/entity"
	errors "github.com/ormushq/ormus/manager/error"
	"github.com/ormushq/ormus/manager/service/auth"

	"github.com/stretchr/testify/assert"
)

func Test_NewJWT(t *testing.T) {
	c := config.C()
	fmt.Println(c)
}

func Test_CreateAccessToken(t *testing.T) {
	testCases := []struct {
		name string
		cfg  *auth.JwtConfig
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
			err:  errors.JwtEmptyUserErr,
		},
		{
			name: "nil user",
			err:  errors.JwtEmptyUserErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			jwt := auth.NewJWT(tc.cfg)

			// 2. execution
			_, err := jwt.CreateAccessToken(tc.user)

			// 3. assertion
			if err != nil {
				assert.Error(t, tc.err, err)
			}
		})
	}
}
