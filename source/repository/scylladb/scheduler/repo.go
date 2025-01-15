package scheduler

import (
	"github.com/ormushq/ormus/source/repository/scylladb"
)

type Repository struct {
	db *scylladb.DB
}

var statements = map[string]scylladb.Statement{}

func New(db *scylladb.DB) *Repository {
	db.RegisterStatements(statements)

	return &Repository{
		db: db,
	}
}
