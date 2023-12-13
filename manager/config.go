package manager

import "github.com/ormushq/ormus/manager/service/auth"

type Config struct {
	JWTConfig auth.JwtConfig `koanf:"jwt_config"`
}
