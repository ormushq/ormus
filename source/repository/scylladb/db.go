package scylladb

import (
	"fmt"
	"log"

	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/adapter/scylladb/scyllainitialize"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/scylladb/gocqlx/v2"
)

type DB struct {
	conn scylladb.SessionxInterface
}

type Statement struct {
	Query  string
	Values []string
}

var statements = map[string]gocqlx.Queryx{}

func New(scylladbConfig scylladb.Config) (*DB, error) {
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
	err = scyllainitialize.RunMigrations(Sconn, "./source/repository/scylladb/")
	if err != nil {
		panic(err)
	}
	Session, Err := scyllainitialize.GetConnection(Sconn)
	if Err != nil {
		panic(Err)
	}

	return &DB{
		conn: Session,
	}, nil
}

func (d *DB) GetConn() scylladb.SessionxInterface {
	return d.conn
}

func (d *DB) RegisterStatement(states ...Statement) {
	for _, stat := range states {
		logger.L().Debug(fmt.Sprintf("%+v", stat))
		statements[stat.Query] = d.conn.Query(stat.Query, stat.Values)
	}
}

func (d *DB) GetStatement(state Statement) (gocqlx.Queryx, error) {
	if statement, ok := statements[state.Query]; ok {
		return statement, nil
	}

	return gocqlx.Queryx{}, richerror.New("db.GetStatement").WhitKind(richerror.KindNotFound).WithMessage("statement not found")
}

func (d *DB) RegisterStatements(states map[string]Statement) {
	for _, stat := range states {
		d.RegisterStatement(stat)
	}
}
