/*
Package initializedb provides functions for initializing a ScyllaDB connection and obtaining a database session.

Usage:

	func main() {
	    // Create a new ScyllaDB connection instance
	    connection := initializedb.NewScyllaDBConnection(gocql.Quorum, "example_keyspace", "127.0.0.1")

	    // Get a ScyllaDB session using the created connection
	    session, err := initializedb.GetConnection(connection)
	    if err != nil {
	        log.Fatal("Failed to get ScyllaDB session:", err)
	    }

	    // Use the 'session' for database operations

	    // Close the session when done
	    defer session.Close()
	}

This package includes functions for creating a new ScyllaDB connection and obtaining a ScyllaDB session.
It utilizes the gocql library for interacting with ScyllaDB.

Functions:

  - NewScyllaDBConnection: Creates and returns a new instance of the 'scyllaDBConnection' type with the specified connection parameters.
    func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, hosts ...string) *scyllaDBConnection

  - GetConnection: Returns a ScyllaDB session using the provided 'scyllaDBConnection' instance.
    It internally creates a ScyllaDB cluster configuration and session.
    func GetConnection(conn *scyllaDBConnection) (scylladb.SessionxInterface, error)
*/
package scyllainitialize

import (
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
	scyllaMigrate "github.com/ormushq/ormus/source/repository/scylladb/migrate"
)

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, hosts ...string) *ScyllaDBConnection {
	return &ScyllaDBConnection{
		consistency: consistency,
		keyspace:    keyspace,
		hosts:       hosts,
	}
}

func GetConnection(conn *ScyllaDBConnection) (scylladb.SessionxInterface, error) {
	return conn.createSession(conn.createCluster())
}

func CreateKeySpace(consistency gocql.Consistency, keyspace string, hosts ...string) error {
	scyllaDBConnection := &ScyllaDBConnection{
		consistency: consistency,
		keyspace:    "system",
		hosts:       hosts,
	}

	session, err := scyllaDBConnection.createSession(scyllaDBConnection.createCluster())
	if err != nil {
		return err
	}

	return scyllaDBConnection.createKeyspace(session, keyspace)
}

func RunMigrations(dbConn *ScyllaDBConnection, dir string) error {
	for _, host := range dbConn.hosts {
		manager := scyllaMigrate.New(dir, host, dbConn.keyspace)
		err := manager.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
