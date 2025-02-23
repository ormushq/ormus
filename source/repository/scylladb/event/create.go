package event

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/repository/scylladb"
	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	statements["create"] = scylladb.Statement{
		Query: `INSERT INTO event (
            write_key, 
            id, 
            type, 
            name, 
            send_at, 
            received_at, 
            event, 
            timestamp, 
            created_at, 
            updated_at,
            delivered
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`,
		Values: []string{
			"write_key",
			"id",
			"type",
			"name",
			"send_at",
			"received_at",
			"event",
			"timestamp",
			"created_at",
			"updated_at",
			"delivered",
		},
	}
}

func (r Repository) CreateNewEvent(ctx context.Context, evt []event.CoreEvent, wg *sync.WaitGroup, queueName string) ([]string, error) {
	messages := make([]*proto.NewEvent, 0)
	ids := make([]string, 0)
	batch := r.db.NewBatch(ctx)
	stmt, err := r.db.GetStatement(statements["create"])
	if err != nil {
		logger.L().Error(err.Error())

		return nil, richerror.New("source.repository").WithWrappedError(err).WithMessage("failed to get prepared statement")
	}
	for _, e := range evt {
		id := uuid.New().String()
		ids = append(ids, id)
		stmt.BindMap(qb.M{
			"write_key":   e.WriteKey,
			"id":          id,
			"type":        e.Type,
			"name":        e.Name,
			"send_at":     e.SendAt,
			"received_at": e.ReceivedAt,
			"event":       e.Event,
			"timestamp":   e.Timestamp,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"properties":  e.Properties,
			"delivered":   false,
		})
		batch.Query(stmt.Statement(), stmt.Values()...)

		messages = append(messages, &proto.NewEvent{
			Id:         id,
			Name:       e.Name,
			WriteKey:   e.WriteKey,
			Event:      e.Event,
			SendAt:     timestamppb.New(e.SendAt),
			ReceivedAt: timestamppb.New(e.ReceivedAt),
			Timestamp:  timestamppb.New(e.Timestamp),
			Type:       string(e.Type),
			Properties: *(e.Properties),
		},
		)
	}
	logger.L().Info(fmt.Sprintf("event %s has been received", messages))
	err = r.db.ExecuteBatch(batch)
	if err != nil {
		logger.L().Error(err.Error())

		return nil, richerror.New("source.repository").WithWrappedError(err).
			WithMessage("failed to get insert statement").
			WithKind(richerror.KindUnexpected)

	}
	err = r.eventBroker.Publish(ctx, queueName, wg, messages)
	if err != nil {
		return []string{}, richerror.New("source.repository").WithWrappedError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return ids, nil
}
