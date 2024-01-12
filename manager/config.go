package manager

import (
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/pkg/cryption"
)

type Config struct {
	JWTConfig       authservice.JwtConfig `koanf:"jwt_config"`
	CryptionConfing cryption.CryptConfing `koanf:"cryption_config"`
}
