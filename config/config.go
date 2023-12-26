package config

import (
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/source"
)

type Config struct {
	Manager manager.Config `koanf:"manager"`
	Redis   redis.Config   `koanf:"redis"`
	Source  source.Config  `koanf:"source"`
}