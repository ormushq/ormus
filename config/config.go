package config

import "github.com/ormushq/ormus/adapter/redis"
import srccfg "github.com/ormushq/ormus/source/config"

type Config struct {
	Source srccfg.Config `koanf:"source"`
	Redis  redis.Config  `koanf:"redis"`
}