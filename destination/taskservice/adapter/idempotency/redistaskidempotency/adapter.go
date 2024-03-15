package redistaskidempotency

import (
	"context"
	"errors"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/redis/go-redis/v9"
)

func (db DB) SaveTaskStatus(ctx context.Context, taskID string, status taskentity.IntegrationDeliveryStatus) error {
	_, err := db.adapter.Client().Set(ctx, db.prefix+taskID, status.String(), db.expiration).Result()
	if err != nil {
		return err
	}

	return nil
}

func (db DB) GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	value, err := db.adapter.Client().Get(ctx, db.prefix+taskID).Result()

	if errors.Is(err, redis.Nil) {
		return taskentity.NotExecutedTaskStatus, nil
	} else if err != nil {
		return taskentity.InvalidTaskStatus, err
	}

	return taskentity.StringToIntegrationDeliveryStatus(value), nil
}
