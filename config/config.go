package config

import "github.com/ormushq/ormus/adapter/redis"

type Config struct {
	Redis redis.Config `koanf:"redis"`
}
