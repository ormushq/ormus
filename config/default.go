package config

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/service/authservice"
)

func Default() Config {
	var accessExpirationTimeInDay time.Duration = 7
	var refreshExpirationTimeInDay time.Duration = 28

	return Config{
		Scylladb: scylladb.Config{
			Hosts:          []string{"127.0.0.1:9042"},
			Consistency:    gocql.One,
			Keyspace:       "default",
			TimeoutCluster: 5 * time.Second,
			NumRetries:     5,
			MinRetryDelay:  time.Second,
			MaxRetryDelay:  10 * time.Second,
		},
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
