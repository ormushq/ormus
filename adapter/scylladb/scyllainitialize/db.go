/*
Package scyllainitialize provides functionality for initializing a connection to ScyllaDB database using to gocql library.

Note: Make sure to handle errors appropriately when using this package.
*/
package scyllainitialize

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
)

type ScyllaDBConnection struct {
	consistency gocql.Consistency
	keyspace    string
	hosts       []string
}

const (
	timeoutCluster = 5 * time.Second

	// numRetries represents the number of retries in the ExponentialBackoffRetryPolicy.
	numRetries = 5

	// minRetryDelay represents the minimum delay duration in the ExponentialBackoffRetryPolicy.
	minRetryDelay = time.Second

	// maxRetryDelay represents the maximum delay duration in the ExponentialBackoffRetryPolicy.
	maxRetryDelay = 10 * time.Second
)

/*
The 'createCluster' method creates and returns a gocql.ClusterConfig
based on the provided connection parameters. It also sets additional configurations,
such as timeout and retry policies.
*/
func (conn *ScyllaDBConnection) createCluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(conn.hosts...)
	cluster.Consistency = conn.consistency
	cluster.Keyspace = conn.keyspace
	cluster.Timeout = timeoutCluster
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		NumRetries: numRetries,
		Min:        minRetryDelay,
		Max:        maxRetryDelay,
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	log.Println("cluster was created.")

	return cluster
}

func (conn *ScyllaDBConnection) createKeyspace(session scylladb.SessionxInterface, keyspace string) error {
	stmt := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}", keyspace)

	return session.ExecStmt(stmt)
}

/*
The 'createSession' method creates a ScyllaDB session using the given gocql.ClusterConfig.
It returns a session wrapped by the 'scylladb' package, which provides additional functionalities.
If an error occurs during the session creation, an error is returned.
*/
func (conn *ScyllaDBConnection) createSession(cluster *gocql.ClusterConfig) (scylladb.SessionxInterface, error) {
	session, err := scylladb.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Println("an error occureed while creating DB Session", err.Error())

		return nil, err
	}

	log.Println("session was created")

	return session, nil
}
