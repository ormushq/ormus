package scyllarepo

import (
	"log"

	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/adapter/scylladb/scyllainitialize"
)

// TODO: implement repository for authservice

type StorageAdapter struct {
	ScyllaConn scylladb.SessionxInterface
}

func New(scylladbConfig scylladb.Config) (*StorageAdapter, error) {
	cfg := scylladb.Config{
		Hosts:          scylladbConfig.Hosts,
		Consistency:    scylladbConfig.Consistency,
		Keyspace:       scylladbConfig.Keyspace,
		TimeoutCluster: scylladbConfig.TimeoutCluster,
		NumRetries:     scylladbConfig.NumRetries,
		MinRetryDelay:  scylladbConfig.MinRetryDelay,
		MaxRetryDelay:  scylladbConfig.MaxRetryDelay,
	}
	Sconn := scyllainitialize.NewScyllaDBConnection(cfg.Consistency, cfg.Keyspace, cfg.Hosts[0])

	err := scyllainitialize.CreateKeySpace(
		cfg.Consistency,
		cfg.Keyspace,
		cfg.Hosts...,
	)
	if err != nil {
		log.Fatal("Failed to create ScyllaDB keyspace:", err)
	}
	err = scyllainitialize.RunMigrations(Sconn, "../../manager/repository/scyllarepo/")
	if err != nil {
		panic(err)
	}
	Session, Err := scyllainitialize.GetConnection(Sconn)
	if Err != nil {
		panic(Err)
	}

	return &StorageAdapter{
		ScyllaConn: Session,
	}, nil
}
