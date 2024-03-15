package inmemorytaskidempotency

import "github.com/ormushq/ormus/destination/entity/taskentity"

func (db DB) SaveTaskStatus(taskID string, status taskentity.IntegrationDeliveryStatus) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.tasks[taskID] = status

	return nil
}

func (db DB) GetTaskStatusByID(taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	status, ok := db.tasks[taskID]
	if !ok {
		return taskentity.NotExecutedTaskStatus, nil
	}

	return status, nil
}
