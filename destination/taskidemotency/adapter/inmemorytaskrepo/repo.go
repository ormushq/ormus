package inmemorytaskrepo

import "github.com/ormushq/ormus/destination/entity"

func (db DB) UpsertTask(taskID string, status entity.TaskStatus) error {

	db.tasks[taskID] = status

	return nil
}

func (db DB) GetStatusByTaskID(taskID string) (entity.TaskStatus, error) {

	status, ok := db.tasks[taskID]
	if !ok {
		return entity.NOT_EXISTS, nil
	}

	return status, nil
}
