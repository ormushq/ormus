package rediswritekey

import (
	"context"
	"errors"
	"github.com/ormushq/ormus/adapter/redis"
	rds "github.com/redis/go-redis/v9"
)

const (
	WriteKeyValidFlag = "valid"
)

type DB struct {
	adapter redis.Adapter
}

// New is Constructor redis DB.
func New(adapter redis.Adapter) DB {
	return DB{adapter: adapter}
}

func (db DB) IsValidWriteKey(ctx context.Context, writeKey string) (bool, error) {
	val, err := db.adapter.Client().Get(ctx, writeKey).Result()
	if errors.Is(err, rds.Nil) {
		// Write key does not exists
		return false, nil
	} else if err != nil {
		// Some other error took place
		return false, err
	}

	// write key exists
	return val == WriteKeyValidFlag, nil
}
