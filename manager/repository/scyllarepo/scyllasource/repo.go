package scyllasource

import (
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
)

type Repository struct {
	db *scyllarepo.DB
}

var statements = map[string]scyllarepo.Statement{}

func New(db *scyllarepo.DB) *Repository {
	db.RegisterStatements(statements)

	return &Repository{
		db: db,
	}
}
