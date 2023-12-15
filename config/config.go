package config

import (
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/manager"
)

type Config struct {
	Manager manager.Config `koanf:"manager"`
	Redis   redis.Config   `koanf:"redis"`
}
