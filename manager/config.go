package manager

import (
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/manager/service/authservice"
)

type Config struct {
	JWTConfig            authservice.JwtConfig `koanf:"jwt_config"`
	InternalBrokerConfig InternalBrokerConfig  `koanf:"internal_broker_config"`
	ScyllaDBConfig       scylladb.Config       `koanf:"scylla_DB_Config"`
	HTTPAddress          string                `koanf:"http_address"`
	GRPCServiceAddress   string                `koanf:"grpc_service_address"`
}

type InternalBrokerConfig struct {
	ChannelSize    int `koanf:"channel_size"`
	NumberInstant  int `koanf:"number_instant"`
	MaxRetryPolicy int `koanf:"max_retry_policy"`
}
