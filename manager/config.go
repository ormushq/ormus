package manager

import (
	"github.com/ormushq/ormus/manager/service/authservice"
)

type Config struct {
	JWTConfig            authservice.JwtConfig `koanf:"jwt_config"`
	InternalBrokerConfig InternalBrokerConfig  `koanf:"internal_broker_config"`
}
type InternalBrokerConfig struct {
	ChannelSize    int `koanf:"channel_size"`
	NumberInstant  int `koanf:"number_instant"`
	MaxRetryPolicy int `koanf:"max_retry_policy"`
}
