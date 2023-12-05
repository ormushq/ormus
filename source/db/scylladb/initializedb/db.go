/*
Package initializedb provides functionality for initializing a connection to ScyllaDB database using to gocql library.

Note: Make sure to handle errors appropriately when using this package.
*/

package initializedb

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/source/db/scylladb"
	"log"
	"time"
)

type ScyllaDBConnection struct {
	consistency gocql.Consistency
	keyspace    string
	hosts       []string
}

/*
The 'createCluster' method creates and returns a gocql.ClusterConfig
based on the provided connection parameters. It also sets additional configurations,
such as timeout and retry policies.
*/
func (conn *ScyllaDBConnection) createCluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(conn.hosts...)
	cluster.Consistency = conn.consistency
	cluster.Keyspace = conn.keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		NumRetries: 5,
		Min:        time.Second,
		Max:        10 * time.Second,
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	log.Println("cluster was created.")
	return cluster
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
