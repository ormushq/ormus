package inmemorytaskmanager

import (
	"github.com/ormushq/ormus/destination/integrationhandler"
	"log"
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

func (w *Worker) ProcessJobs() {

	var forever chan struct{}
	go func() {

		for task := range w.TaskManager.Queue.Tasks {

			log.Printf("Task [%s] received by In-Memory worker.", task.ID)
			err := w.Handler.Handle(task.ProcessedEvent)
			if err != nil {
				//todo change task status to success
				//w.TaskManager.taskStorage.ChangeTaskStatus("", "FAILED")

			}

			//todo change task status to success
			//w.TaskManager.taskStorage.ChangeTaskStatus("", "SUCCESS")
		}
	}()

	log.Printf(" [In-Memory] Waiting for messages. To exit press CTRL+C")
	//for running workers independently
	<-forever
}
