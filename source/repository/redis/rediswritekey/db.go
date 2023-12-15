package rediswritekey

import "github.com/ormushq/ormus/adapter/redis"

type DB struct {
	adapter redis.Adapter
}

// New is Constructor redis DB.
func New(adapter redis.Adapter) DB {
	return DB{adapter: adapter}
}
