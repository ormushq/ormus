package config

import (
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/source"
)

type Config struct {
	Source source.Config `koanf:"source"`
	Redis  redis.Config  `koanf:"redis"`
}