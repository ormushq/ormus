package redistaskidempotency

import (
	"context"
	"errors"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (db DB) SaveTaskStatus(ctx context.Context, taskID string, status taskentity.IntegrationDeliveryStatus) error {
	tracer := otela.NewTracer("redistaskidempotency")
	_, span := tracer.Start(ctx, "redistaskidempotency@SaveTaskStatus", trace.WithAttributes(
		attribute.String("task_id", taskID),
		attribute.String("status", status.String()),
	))
	defer span.End()

	_, err := db.adapter.Client().Set(ctx, db.prefix+taskID, status.ToNumericString(), db.expiration).Result()
	if err != nil {
		return err
	}
	span.AddEvent("task-status-added-to-db")

	return nil
}

func (db DB) GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	tracer := otela.NewTracer("redistaskidempotency")
	ctx, span := tracer.Start(ctx, "redistaskidempotency@GetTaskStatusByID", trace.WithAttributes(
		attribute.String("task_id", taskID),
	))
	defer span.End()

	value, err := db.adapter.Client().Get(ctx, db.prefix+taskID).Result()

	if errors.Is(err, redis.Nil) {
		span.AddEvent("task-not-found-on-db")

		return taskentity.NotExecutedTaskStatus, nil
	} else if err != nil {
		span.AddEvent("error-on-fetch-task", trace.WithAttributes(
			attribute.String("error", err.Error())))

		return taskentity.InvalidTaskStatus, err
	}
	result := taskentity.NumericStringToIntegrationDeliveryStatus(value)
	span.AddEvent("task-status-fetched", trace.WithAttributes(
		attribute.String("status", result.String())))

	return result, nil
}
