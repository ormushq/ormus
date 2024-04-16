package rabbitmqtaskmanager

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler"
	"github.com/ormushq/ormus/destination/taskmanager"
	"github.com/ormushq/ormus/destination/taskservice"
	"github.com/ormushq/ormus/event"
)

const timeoutInSeconds = 5

type Worker struct {
	TaskConsumer taskmanager.Consumer
	Handler      integrationhandler.IntegrationHandler
	TaskService  taskservice.Service
}

func (w *Worker) handleEvent(ctx context.Context, newEvent event.ProcessedEvent) error {
	var task taskentity.Task
	var err error

	ts := w.TaskService
	var taskStatus taskentity.IntegrationDeliveryStatus

	taskID := newEvent.ID()

	// check idempotency
	if taskStatus, err = ts.GetTaskStatusByID(ctx, taskID); err != nil {
		// todo use richError
		return err
	}

	if !taskStatus.CanBeExecuted() {
		slog.Debug(fmt.Sprintf("Task [%s] has %s status and is not executable", taskID, taskStatus.String()))

		return nil
	}

	if taskStatus.IsBroadcast() {
		// Get all task info (attempts, failed reason and...) from repository.
		task, err = ts.GetTaskByID(taskID)
		if err != nil {
			return err
		}
	} else {
		task = taskentity.MakeTaskUsingProcessedEvent(newEvent)
	}

	res, err := w.Handler.Handle(task, newEvent)
	if err != nil {
		return err
	}

	task.IntegrationDeliveryStatus = res.DeliveryStatus
	task.Attempts = res.Attempts
	task.FailedReason = res.FailedReason

	err = ts.UpsertTaskAndSaveIdempotency(ctx, task)
	if err != nil {
		// todo what should we do if error occurs in updating task repo or idempotency ?
		return err
	}

	return nil
}

func (w *Worker) Run(done <-chan bool, wg *sync.WaitGroup) error {
	processedEventsChannel, err := w.TaskConsumer.Consume(done, wg)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Start rabbitmq worker for handling tasks.")

		for {
			select {
			case newEvent := <-processedEventsChannel:
				go func() {
					ctx, cancel := context.WithTimeout(context.Background(), timeoutInSeconds*time.Second)
					defer cancel()
					err := w.handleEvent(ctx, newEvent)
					if err != nil {
						slog.Error(fmt.Sprintf("Error on handling event using integration handler.Error : %v", err))
					}
				}()
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
