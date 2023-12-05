package initializedb

import (
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/source/db/scylladb"
)

func NewScyllaDBConnection(consistency gocql.Consistency, keyspace string, hosts ...string) *scyllaDBConnection {
	return &scyllaDBConnection{
		consistency: consistency,
		keyspace:    keyspace,
		hosts:       hosts,
	}
}

func GetConnection(conn *scyllaDBConnection) (scylladb.SessionxInterface, error) {
	return conn.createSession(conn.createCluster())
}
