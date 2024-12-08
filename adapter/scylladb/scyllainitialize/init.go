package scyllainitialize

import (
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/logger"
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
	logger.L().Debug("running migrations...")
	for _, host := range dbConn.hosts {
		migration := New(dir, host, dbConn.keyspace)
		err := migration.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
