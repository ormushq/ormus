package inmemorytaskmanager

import (
	"github.com/ormushq/ormus/destination/entity"
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

		for task := range w.TaskManager.Queue.tasks {

			ti := w.TaskManager.taskIdempotency
			var taskStatus entity.TaskStatus
			taskID := task.ID

			log.Printf("Task [%s] received by In-Memory worker.", taskID)

			enabled, err := ti.IntegrationHandlerIsEnable(taskID)

			if err != nil {
				log.Println("Error on IntegrationHandlerIsEnable.", err)
				continue
			}

			if enabled {
				err := w.Handler.Handle(task.ProcessedEvent)
				if err != nil {
					taskStatus = entity.FAILED_IN_INTEGRATION_HANDLER
				}
			} else {
				log.Printf("\033[33mPrevent to handling duplicate processed event in idempotency.!\033[0m\n")

			}

			taskStatus = entity.SUCCESS_IN_INTEGRATION_HANDLER

			err = ti.Save(taskID, taskStatus)
			if err != nil {
				log.Println("Error on Saving task status.", err)
				continue
			}
		}
	}()

	log.Printf(" [In-Memory] Waiting for messages. To exit press CTRL+C")
	//for running workers independently
	<-forever
}
