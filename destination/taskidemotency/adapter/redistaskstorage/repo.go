package redistaskstorage

import "github.com/ormushq/ormus/destination/entity"

func (db DB) UpsertTask(taskID string, status entity.TaskStatus) error {

	//todo implement redis UpsertTask

	return nil
}

func (db DB) GetStatusByTaskID(taskID string) (entity.TaskStatus, error) {

	//todo implement redis GetStatusByTaskID

	return entity.NOT_EXISTS, nil
}
