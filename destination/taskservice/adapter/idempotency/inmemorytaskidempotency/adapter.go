package inmemorytaskidempotency

import (
	"context"

	"github.com/ormushq/ormus/destination/entity/taskentity"
)

func (db DB) SaveTaskStatus(_ context.Context, taskID string, status taskentity.IntegrationDeliveryStatus) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.tasks[taskID] = status

	return nil
}

func (db DB) GetTaskStatusByID(_ context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	status, ok := db.tasks[taskID]
	if !ok {
		return taskentity.NotExecutedTaskStatus, nil
	}

	return status, nil
}
