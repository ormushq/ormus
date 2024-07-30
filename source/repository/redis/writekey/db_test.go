package writekey_test

import (
	"google.golang.org/grpc"
	"os"
	"testing"

	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/repository/redis/writekey"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (writekey.DB, func()) {
	redisAdapter, err := redis.New(config.C().Redis)
	assert.Nil(t, err)

	managerClientConn, err := grpc.Dial(config.C().Manager.GRPCServiceAddress, grpc.WithInsecure())
	assert.Nil(t, err)

	redisRepository := writekey.New(redisAdapter, managerClientConn)
	return redisRepository, func() {
		err := managerClientConn.Close()
		if err != nil {
			return
		}
	}
}
