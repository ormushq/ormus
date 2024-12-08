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
	HTTPPort int `koanf:"http_port"`
	GrpcPort int `koanf:"grpc_port"`
}

type InternalBrokerConfig struct {
	ChannelSize    int `koanf:"channel_size"`
	MaxRetryPolicy int `koanf:"max_retry_policy"`
}
