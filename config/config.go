package config

import "github.com/ormushq/ormus/manager"
import "github.com/ormushq/ormus/adapter/redis"

type Config struct {
  Manager manager.Config `koanf:"manager"`
	Redis redis.Config `koanf:"redis"`
}
