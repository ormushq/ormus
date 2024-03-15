package rabbitmqtaskmanager

import (
	"log"
	"sync"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler"
	"github.com/ormushq/ormus/destination/taskmanager"
	"github.com/ormushq/ormus/destination/taskservice"
)

type Worker struct {
	TaskConsumer taskmanager.Consumer
	Handler      integrationhandler.IntegrationHandler
	TaskService  taskservice.Service
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) error {
	processedEventsChannel, err := w.TaskConsumer.Consume(done, wg)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting rabbitmq worker...")

		for {
			select {
			case newEvent := <-processedEventsChannel:

				var task *taskentity.Task
				var err error

				ts := w.TaskService
				var taskStatus taskentity.IntegrationDeliveryStatus
				var failedReason *string

				// todo get ID using function
				taskID := newEvent.ID()

				// check idempotency
				if taskStatus, err = ts.GetTaskStatusByID(taskID); err != nil {
					// todo what should we do if error occurs in idempotency ?

					log.Println("Error on GetTaskStatusByID.", err)
				}

				if taskStatus.CanBeExecuted() {

					if taskStatus.IsBroadcast() {
						task, err = ts.GetTaskByID(taskID)
						if err != nil {
							log.Println("Error on GetTaskByID.", err)
						}
					} else {
						task = taskentity.MakeTaskUsingProcessedEvent(newEvent)
					}

					res, err := w.Handler.Handle(task)
					if err != nil {
						failedReason = res.ErrorReason
						taskStatus = res.DeliveryStatus
					} else {
						taskStatus = taskentity.Success
					}
					err = ts.UpsertTaskAndSaveIdempotency(*task, taskStatus, failedReason)
					if err != nil {
						// todo what should we do if error occurs in updating task storage ?
						// todo what should we do if error occurs in updating idempotency ?
						// todo should we consider this function as an atomic operation ?
						log.Println("Error on UpsertTaskAndSaveIdempotency.", err)
					}

				} else {
					log.Printf("Task [%s] is not executable", taskID)
				}

			case <-done:

				return
			}
		}
	}()

	return nil
}

func NewWorker(c taskmanager.Consumer, h integrationhandler.IntegrationHandler, srv taskservice.Service) *Worker {
	return &Worker{
		TaskConsumer: c,
		Handler:      h,
		TaskService:  srv,
	}
}

func panicOnWorkersError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func printWorkersError(err error, msg string) {
	log.Printf("%s: %s", msg, err)
}
