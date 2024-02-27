package taskmanager

import "github.com/ormushq/ormus/destination/entity"

type TaskManager interface {
	publish(task entity.Task) error
}
