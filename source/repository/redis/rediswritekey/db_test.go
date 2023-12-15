package rediswritekey_test

import (
	"os"
	"testing"

	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/repository/redis/rediswritekey"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (rediswritekey.DB, func()) {
	redisAdapter, err := redis.New(config.C().Redis)
	assert.Nil(t, err)
	redisRepository := rediswritekey.New(redisAdapter)
	return redisRepository, func() {
		// TODO - cleanup
	}
}
