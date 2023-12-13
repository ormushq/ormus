package config

import srccfg "github.com/ormushq/ormus/source/config"

type Config struct {
	Source srccfg.Config `koanf:"source"`
}
