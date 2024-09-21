package config

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/service/authservice"
)

const (
	TimeoutCluster = 5 * time.Second

	// NumRetries represents the number of retries in the ExponentialBackoffRetryPolicy.
	NumRetries = 5

	// MinRetryDelay represents the minimum delay duration in the ExponentialBackoffRetryPolicy.
	MinRetryDelay = time.Second

	// MaxRetryDelay represents the maximum delay duration in the ExponentialBackoffRetryPolicy.
	MaxRetryDelay = 10 * time.Second
)

const (
	ChannelSize    = 100
	NumberInstant  = 100
	MaxRetryPolicy = 10
)

func Default() Config {
	accessExpirationTimeInDay := 7
	refreshExpirationTimeInDay := 28

	return Config{
		Scylladb: scylladb.Config{
			Hosts:          []string{"127.0.0.1:9042"},
			Consistency:    gocql.One,
			Keyspace:       "default",
			TimeoutCluster: TimeoutCluster,
			NumRetries:     NumRetries,
			MinRetryDelay:  MinRetryDelay,
			MaxRetryDelay:  MaxRetryDelay,
		},
		Manager: manager.Config{
			AuthConfig: authservice.Config{
				SecretKey:                  "Ormus_jwt",
				AccessExpirationTimeInDay:  accessExpirationTimeInDay,
				RefreshExpirationTimeInDay: refreshExpirationTimeInDay,
				AccessSubject:              "ac",
				RefreshSubject:             "rt",
			},
			InternalBrokerConfig: manager.InternalBrokerConfig{
				ChannelSize:    ChannelSize,
				NumberInstant:  NumberInstant,
				MaxRetryPolicy: MaxRetryPolicy,
			},
		},
	}
}
