package config

import (
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

func Load(configPath string) Config {
	// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	var k = koanf.New(".")

	// Load environment variables and merge into the loaded config.
	// "MYVAR" is the prefix to filter the env vars by.
	// "." is the delimiter used to represent the key hierarchy in env vars.
	// The (optional, or can be nil) function can be used to transform
	// the env var names, for instance, to lowercase them.
	//
	// For `example`, env vars: MYVAR_TYPE and MYVAR_PARENT1_CHILD1_NAME
	// will be merged into the "type" and the nested "parent1.child1.name"
	// keys in the config file here as we lowercase the key,
	// replace `_` with `.` and strip the MYVAR_ prefix so that
	// only "parent1.child1.name" remains.

	err := k.Load(file.Provider(configPath), yaml.Parser()); if err != nil {
		panic(err)
	}

	err = k.Load(env.Provider("ORMUS_SOURCE_", ".", func(s string) string {
		
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, "ORMUS_SOURCE_")), "_", ".")
	}), nil); if err != nil {
		panic(err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return cfg
}
