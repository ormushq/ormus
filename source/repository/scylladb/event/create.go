package event

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/repository/scylladb"
	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	statements["create"] = scylladb.Statement{
		Query:  `INSERT INTO event (id, type, name, send_at, received_at, timestamp, event, write_key, created_at, updated_at, properties,delivered) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,false)`,
		Values: []string{"id", "type", "name", "send_at", "received_at", "timestamp", "event", "write_key", "created_at", "updated_at", "properties"},
	}
}

func (r Repository) CreateNewEvent(ctx context.Context, evt event.CoreEvent, wg *sync.WaitGroup, queueName string) (string, error) {
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
	messages := []*proto.NewEvent{
		{
			Id:         id,
			Name:       evt.Name,
			WriteKey:   evt.WriteKey,
			Event:      evt.Event,
			SendAt:     timestamppb.New(evt.SendAt),
			ReceivedAt: timestamppb.New(evt.ReceivedAt),
			Timestamp:  timestamppb.New(evt.Timestamp),
			Type:       string(evt.Type),
			Properties: *(evt.Properties),
		},
	}
	err = r.eventBroker.Publish(ctx, queueName, wg, messages)
	if err != nil {
		return "", richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return id, nil
}
