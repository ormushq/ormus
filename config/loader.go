package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

const (
	prefix    = "ORMUS_"
	delimiter = "."
	separator = "__"
)

// our environment variables must prefix with `ORMUS_`
// for nested env should use `__` aka: ORMUS_DB__HOST.
func callbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, prefix))

	return strings.ReplaceAll(base, separator, delimiter)
}

func Load() Config {
	// dot means the delimiter
	k := koanf.New(".")

	// load default configuration from Default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default config: %s", err)
	}

	// load configuration from yaml file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config from `config.yml` file: %s", err)
	}

	// load from environment variable
	if err := k.Load(env.Provider(prefix, delimiter, callbackEnv), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	var instance Config
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	fmt.Printf("%+v\n", instance)

	return instance
}
