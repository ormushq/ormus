package event

import (
	"context"

	"github.com/ormushq/ormus/contract/go/destination"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/repository/scylladb"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["delivered_event"] = scylladb.Statement{
		Query:  `UPDATE event SET delivered = true WHERE  id = ?`,
		Values: []string{"id"},
	}
}

func (r Repository) EventHasDelivered(ctx context.Context, evt *destination.DeliveredEventsList) error {
	query, err := r.db.GetStatement(statements["delivered_event"])
	batch := r.db.NewBatch(ctx)
	if err != nil {
		logger.L().Error(err.Error())
		return richerror.New("source.repository").WithWrappedError(err).
			WithMessage("failed to  update delivered status due to delivered_event statement").
			WithKind(richerror.KindUnexpected)

	}
	for _, e := range evt.Events {
		query.BindMap(qb.M{
			"id": e.MessageId,
		})
		batch.Query(query.Statement(), query.Values()...)
	}
	err = r.db.ExecuteBatch(batch)
	if err != nil {
		logger.L().Error(err.Error())
		return richerror.New("source.repository").WithWrappedError(err).
			WithMessage("failed to update delivered status due to bulk update failure").
			WithKind(richerror.KindUnexpected)

	}
	return nil
}
