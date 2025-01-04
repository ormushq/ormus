package event

import (
	"github.com/ormushq/ormus/source/eventhandler"
	"github.com/ormushq/ormus/source/repository/scylladb"
)

type Repository struct {
	db          *scylladb.DB
	eventBroker eventhandler.Publisher
}

var statements = map[string]scylladb.Statement{}

func New(db *scylladb.DB, publisher eventhandler.Publisher) *Repository {
	db.RegisterStatements(statements)

	return &Repository{
		db:          db,
		eventBroker: publisher,
	}
}
