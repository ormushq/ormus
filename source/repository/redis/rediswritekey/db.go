package rediswritekey

import (
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/source/adapter/manager"
	"golang.org/x/sync/singleflight"
)

type DB struct {
	adapter        redis.Adapter
	ManagerAdapter manager.Manager
	singleFlight   singleflight.Group
}

// New is Constructor redis DB.
func New(adapter redis.Adapter, managerAdapter manager.Manager) DB {
	return DB{
		adapter:        adapter,
		ManagerAdapter: managerAdapter,
	}
}
