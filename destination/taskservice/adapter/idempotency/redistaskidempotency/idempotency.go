package redistaskidempotency

import (
	"fmt"

	"github.com/ormushq/ormus/destination/entity/taskentity"
)

func (db DB) SaveTaskStatus(taskID string, status taskentity.IntegrationDeliveryStatus) error {
	// todo implement redis idempotency

	fmt.Println("Task ID:", taskID, "Status:", status)

	return nil
}

func (db DB) GetTaskStatusByID(taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	// todo implement redis idempotency
	fmt.Println("Task ID:", taskID, "Status:")

	return taskentity.InvalidTaskStatus, nil
}
