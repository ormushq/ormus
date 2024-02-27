package inmemorytaskrepo

import "github.com/ormushq/ormus/destination/entity"

type DB struct {
	tasks map[string]entity.TaskStatus
}

func New() DB {
	tasks := make(map[string]entity.TaskStatus)
	return DB{tasks: tasks}
}
