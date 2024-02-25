package taskmanager

import "github.com/ormushq/ormus/destination/entity"

type TaskManager interface {
	SendToQueue(task entity.Task) error
}
