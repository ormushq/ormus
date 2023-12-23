/*
Package scylladb provides functionality for initializing a connection to ScyllaDB database using to gocql library.

Note: Make sure to handle errors appropriately when using this package.
*/
package scylladb

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type Config struct {
	Hosts          []string
	Consistency    gocql.Consistency
	Keyspace       string
	TimeoutCluster time.Duration
	NumRetries     int
	MinRetryDelay  time.Duration
	MaxRetryDelay  time.Duration
}

const (
	TimeoutCluster = 5 * time.Second

	// NumRetries represents the number of retries in the ExponentialBackoffRetryPolicy.
	NumRetries = 5

	// MinRetryDelay represents the minimum delay duration in the ExponentialBackoffRetryPolicy.
	MinRetryDelay = time.Second

	// MaxRetryDelay represents the maximum delay duration in the ExponentialBackoffRetryPolicy.
	MaxRetryDelay = 10 * time.Second
)

// New
// The package defines a configuration structure (Config)
// and a function (New) to create and return a session interface for
// interacting with ScyllaDB. The New function takes a Config parameter
// to customize the behavior of the ScyllaDB connection.
//
//		func main() {
//			// Set up ScyllaDB configuration
//			config := scylladb.Config{
//				Hosts:          []string{"127.0.0.1:9042"},
//				Consistency:    gocql.One,
//				Keyspace:       "your_keyspace",
//				TimeoutCluster: 5 * time.Second,
//				NumRetries:     5,
//				MinRetryDelay:  time.Second,
//				MaxRetryDelay:  10 * time.Second,
//			}
//
//			// Create a new ScyllaDB session
//			session, err := scylladb.New(config)
//			if err != nil {
//				log.Fatal("Error creating ScyllaDB session:", err)
//			}
//	     defer session.Close()
//
//			// Use the session to interact with ScyllaDB
//
//		}
//
// /**
func New(config Config) (SessionxInterface, error) {
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Consistency = config.Consistency
	cluster.Keyspace = config.Keyspace
	cluster.Timeout = config.TimeoutCluster
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		NumRetries: config.NumRetries,
		Min:        config.MinRetryDelay,
		Max:        config.MaxRetryDelay,
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	session, err := WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Println("an error occureed while creating DB Session", err.Error())

		return nil, err
	}

	return session, nil
}
