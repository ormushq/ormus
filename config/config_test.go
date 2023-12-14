package config_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type structConfig struct {
	Debug        bool   `koanf:"debug"`
	MultiWordVar string `koanf:"multi_word_var"`
	DB           db     `koanf:"db"`
}

type db struct {
	Host                  string                `koanf:"host"`
	Username              string                `koanf:"username"`
	Password              string                `koanf:"password"`
	MultiWordNestedVar    string                `koanf:"multi_word_nested_var"`
	NestedMultiWordConfig nestedMultiWordConfig `koanf:"nested_multi_word_config"`
}

type nestedMultiWordConfig struct {
	DownHere string `koanf:"down_here"`
}

const (
	prefix    = "ORMUS_"
	delimiter = "."
	separator = "__"
)

func callbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, prefix))

	return strings.ReplaceAll(base, separator, delimiter)
}

func TestLoadingDefaultConfigFromStruct(t *testing.T) {
	k := koanf.New(delimiter)

	testStruct := structConfig{
		Debug:        false,
		MultiWordVar: "Im complex in default",
		DB: db{
			Host:               "localhost",
			Username:           "hossein",
			Password:           "1234",
			MultiWordNestedVar: "Oh this is too long",
		},
	}

	if err := k.Load(structs.Provider(testStruct, "koanf"), nil); err != nil {
		t.Fatalf("errmsg loading default config: %s", err)
	}

	var instance structConfig
	if err := k.Unmarshal("", &instance); err != nil {
		t.Fatalf("errmsg unmarshaling config: %s", err)
	}

	if !reflect.DeepEqual(instance, testStruct) {
		t.Fatalf("expected: %+v, got: %+v", testStruct, instance)
	}
}

func TestLoadingConfigFromYamlFile(t *testing.T) {
	k := koanf.New(delimiter)

	ymlConfigTest := []byte(`debug: false
multi_word_var: "I'm complex in config.yml"
db:
  host: "localhost"
  username: "ali"
  password: "passwd"
  multi_word_nested_var: "WHAT??"`)

	ymlFile, _ := os.Create("test.yml")
	defer ymlFile.Close()
	defer os.Remove("test.yml")
	ymlFile.Write(ymlConfigTest)
	// load configuration from yaml file
	if err := k.Load(file.Provider("test.yml"), yaml.Parser()); err != nil {
		t.Logf("errmsg loading config from `test.yml` file: %s", err)
	}

	want := structConfig{
		Debug:        false,
		MultiWordVar: "I'm complex in config.yml",
		DB: db{
			Host:               "localhost",
			Username:           "ali",
			Password:           "passwd",
			MultiWordNestedVar: "WHAT??",
		},
	}

	var instance structConfig
	if err := k.Unmarshal("", &instance); err != nil {
		t.Fatalf("errmsg unmarshaling config: %s", err)
	}

	if !reflect.DeepEqual(want, instance) {
		t.Fatalf("expected: %+v, got: %+v", want, instance)
	}
}

func TestLoadConfigFromEnvironmentVariable(t *testing.T) {
	k := koanf.New(".")

	os.Setenv("ORMUS_DEBUG", "false")
	os.Setenv("ORMUS_MULTI_WORD_VAR", "this is multi word var")
	os.Setenv("ORMUS_DB__HOST", "localhost")
	os.Setenv("ORMUS_DB__USERNAME", "hossein")
	os.Setenv("ORMUS_DB__PASSWORD", "1234")
	os.Setenv("ORMUS_DB__MULTI_WORD_NESTED_VAR", "testing make it easy (:")
	os.Setenv("ORMUS_DB__NESTED_MULTI_WORD_CONFIG__DOWN_HERE", "im here")

	if err := k.Load(env.Provider(prefix, delimiter, callbackEnv), nil); err != nil {
		t.Logf("errmsg loading environment variables: %s", err)
	}

	var instance structConfig
	if err := k.Unmarshal("", &instance); err != nil {
		t.Fatalf("errmsg unmarshaling config: %s", err)
	}

	want := structConfig{
		Debug:        false,
		MultiWordVar: "this is multi word var",
		DB: db{
			Host:                  "localhost",
			Username:              "hossein",
			Password:              "1234",
			MultiWordNestedVar:    "testing make it easy (:",
			NestedMultiWordConfig: nestedMultiWordConfig{DownHere: "im here"},
		},
	}
	if !reflect.DeepEqual(instance, want) {
		t.Fatalf("expected: %+v, got: %+v", want, instance)
	}
}
