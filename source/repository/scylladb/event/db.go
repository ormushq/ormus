package event

import (
	"github.com/ormushq/ormus/adapter/scylladb"
)

type DB struct {
	adapter scylladb.SessionxInterface
}

func New(adapter scylladb.SessionxInterface) DB {
	return DB{adapter: adapter}
}
