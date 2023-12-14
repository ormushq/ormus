package manager

import "github.com/ormushq/ormus/manager/service/auth"

type Config struct {
	JWTConfig service.JwtConfig `koanf:"jwt_config"`
}
