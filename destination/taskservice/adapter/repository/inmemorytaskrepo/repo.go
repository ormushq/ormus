package inmemorytaskrepo

import (
	"fmt"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskservice/param"
)

func (db DB) GetTaskByID(taskID string) (taskentity.Task, error) {
	task, ok := db.tasks[taskID]
	if !ok {
		return taskentity.Task{}, fmt.Errorf("task not found")
	}

	return *task, nil
}

func (db DB) UpsertTask(taskID string, request param.UpsertTaskRequest) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	task, ok := db.tasks[taskID]
	if !ok {
		// todo create task
		task = &taskentity.Task{
			ID: taskID,
		}
		db.tasks[taskID] = task
	}
	task.IntegrationDeliveryStatus = request.IntegrationDeliveryStatus
	task.Attempts = request.Attempts
	task.FailedReason = request.FailedReason

	return nil
}
