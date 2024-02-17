package manager

import (
	"github.com/ormushq/ormus/manager/service/authservice"
)

type Config struct {
	JWTConfig authservice.JwtConfig `koanf:"jwt_config"`
}
