package password_test

import (
	"github.com/ormushq/ormus/pkg/password"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		pass := "very_strong_password"

		// 2. execution
		hashPassword, err := password.HashPassword(pass)

		// 3. assertion
		assert.NoError(t, err)
		assert.NotEmpty(t, hashPassword)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("ordinary", func(t *testing.T) {
		// 1. setup
		pass := "very_strong_password"
		hash, _ := password.HashPassword(pass)

		// 2. execution
		res := password.CheckPasswordHash(pass, hash)

		// 3. assertion
		assert.True(t, res)
	})
}
