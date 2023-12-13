package config

import "github.com/ormushq/ormus/manager"

type Config struct {
	Manager manager.Config `koanf:"manager"`
}
