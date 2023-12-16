package manager

import service "github.com/ormushq/ormus/manager/service/authservice"

type Config struct {
	JWTConfig service.JwtConfig `koanf:"jwt_config"`
}
