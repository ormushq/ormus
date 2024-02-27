package inmemorytaskmanager

import (
	"fmt"
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/destination/taskidemotency/service/taskidempotencyservice"
)

type TaskManager struct {
	Queue           *Queue
	taskIdempotency taskidempotencyservice.Service
}

func NewTaskManager(ti taskidempotencyservice.Service) *TaskManager {
	q := NewQueue()
	return &TaskManager{
		Queue:           q,
		taskIdempotency: ti,
	}
}

func (tm *TaskManager) SendToQueue(t *entity.Task) error {

	err := tm.Queue.Enqueue(t)
	if err != nil {
		fmt.Println("enqueue Error : ", err)
		return err
	}

	return nil
}
