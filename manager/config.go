package manager

import (
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/manager/service/authservice"
)

type Config struct {
	Application          ApplicationConfig    `koanf:"application"`
	AuthConfig           authservice.Config   `koanf:"auth_config"`
	InternalBrokerConfig InternalBrokerConfig `koanf:"internal_broker_config"`
	ScyllaDBConfig       scylladb.Config      `koanf:"scylla_db_config"`
}

type ApplicationConfig struct {
	Port int `koanf:"port"`
}

type InternalBrokerConfig struct {
	ChannelSize    int `koanf:"channel_size"`
	NumberInstant  int `koanf:"number_instant"`
	MaxRetryPolicy int `koanf:"max_retry_policy"`
}
