package config

import (
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/service/auth"
)

func Default() Config {
	return Config{
		Manager: manager.Config{
			JWTConfig: auth.JwtConfig{
				SecretKey:                  "Ormus_jwt",
				AccessExpirationTimeInDay:  7,
				RefreshExpirationTimeInDay: 28,
				AccessSubject:              "ac",
				RefreshSubject:             "rt",
			},
		},
	}
}
