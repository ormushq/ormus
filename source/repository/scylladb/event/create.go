package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/repository/scylladb"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["create"] = scylladb.Statement{
		Query:  `INSERT INTO event (id, type, name, send_at, received_at, timestamp, event, write_key, created_at, updated_at, properties) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		Values: []string{"id", "type", "name", "send_at", "received_at", "timestamp", "event", "write_key", "created_at", "updated_at", "properties"},
	}
}

func (r Repository) CreateNewEvent(evt event.CoreEvent) (string, error) {
	query, err := r.db.GetStatement(statements["create"])
	if err != nil {
		return "", richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	id := uuid.New().String()
	query.BindMap(qb.M{
		"write_key":   evt.WriteKey,
		"id":          id,
		"type":        evt.Type,
		"name":        evt.Name,
		"send_at":     evt.SendAt,
		"received_at": evt.ReceivedAt,
		"event":       evt.Event,
		"timestamp":   evt.Timestamp,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"properties":  evt.Properties,
	})

	if err := query.Exec(); err != nil {
		return "", richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return id, nil
}
