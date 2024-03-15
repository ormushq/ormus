package inmemorytaskmanager

import (
	"log"

	"github.com/ormushq/ormus/destination/integrationhandler"
)

type Worker struct {
	TaskManager *TaskManager
	Handler     integrationhandler.IntegrationHandler
}

func NewWorker(tm *TaskManager, h integrationhandler.IntegrationHandler) *Worker {
	return &Worker{
		TaskManager: tm,
		Handler:     h,
	}
}

func (w *Worker) ExecuteTasks() {
	var forever chan struct{}

	log.Printf(" [In-Memory] Waiting for messages. To exit press CTRL+C")
	// for running workers independently
	<-forever
}
