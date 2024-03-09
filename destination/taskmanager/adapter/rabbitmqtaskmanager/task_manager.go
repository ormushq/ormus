package rabbitmqtaskmanager

import (
	"encoding/json"
	"fmt"

	"github.com/ormushq/ormus/destination/config"
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/destination/taskidemotency/service/taskidempotencyservice"
)

type TaskManager struct {
	queueName       string
	config          config.RabbitMQTaskManagerConnection
	taskIdempotency taskidempotencyservice.Service
	queue           *Queue
}

func NewTaskManager(c config.RabbitMQTaskManagerConnection, queueName string, ti taskidempotencyservice.Service) *TaskManager {
	return &TaskManager{
		queueName:       queueName,
		config:          c,
		taskIdempotency: ti,
	}
}

func (tm *TaskManager) SendToQueue(t *entity.Task) error {
	// do we need a pool for keeping queues ?
	q := NewQueue(tm.config, tm.queueName)

	err := q.Enqueue(t)
	if err != nil {
		// handle enqueue error
	}

	return nil
}

func (tm *TaskManager) UnmarshalMessageToTask(msg []byte) (*entity.Task, error) {
	// todo we can unmarshal base on config (json, protobuf and ...). but for now we assume there is only json option.

	var task entity.Task
	if err := json.Unmarshal(msg, &task); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &task, nil
}
