package rediswritekey

import (
	"context"
	adapter "github.com/ormushq/ormus/adapter/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockRedisAdapter struct {
	mock.Mock
}

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return args.Get(0).(*redis.StatusCmd)
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	return args.Get(0).(*redis.StringCmd)
}

func (m *MockRedisAdapter) Client() *adapter.Adapter {
	args := m.Called()
	return args.Get(0).(*adapter.Adapter)
}

//func TestCreateNewWriteKey_Success(t *testing.T) {
//	mockClient := new(MockRedisClient)
//	mockAdapter := new(MockRedisAdapter)
//
//	mockAdapter.On("Client").Return(mockClient)
//
//	writeKey := "writeKey123"
//	expirationTime := uint(10)
//	mockClient.On("Set", mock.Anything, writeKey, mock.Anything, time.Minute*time.Duration(expirationTime)).
//		Return(&redis.StatusCmd{}).Once() // Mock successful Set command
//
//	db := New(mockAdapter)
//
//}
