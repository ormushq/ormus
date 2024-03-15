package redistaskidempotency

import (
	"time"

	"github.com/ormushq/ormus/adapter/redis"
)

type DB struct {
	adapter    redis.Adapter
	prefix     string
	expiration time.Duration
}

func New(adapter redis.Adapter, prefix string, exp time.Duration) DB {
	return DB{
		adapter:    adapter,
		prefix:     prefix,
		expiration: exp,
	}
}
