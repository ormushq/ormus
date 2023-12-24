package config

import (
	"time"

	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/service/authservice"
)

func Default() Config {
	var accessExpirationTimeInDay time.Duration = 7
	var refreshExpirationTimeInDay time.Duration = 28

	return Config{
		Manager: manager.Config{
			JWTConfig: authservice.JwtConfig{
				SecretKey:                  "Ormus_jwt",
				AccessExpirationTimeInDay:  accessExpirationTimeInDay,
				RefreshExpirationTimeInDay: refreshExpirationTimeInDay,
				AccessSubject:              "ac",
				RefreshSubject:             "rt",
			},
		},
	}
}
