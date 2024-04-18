package taskservice

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery"
	"github.com/ormushq/ormus/event"
)

type TaskService interface {
	GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error)
	GetTaskByID(taskID string) (taskentity.Task, error)
	UpsertTaskAndSaveIdempotency(ctx context.Context, t taskentity.Task) error
}

type Handler struct {
	DeliveryHandler taskdelivery.DeliveryHandler
	TaskService     TaskService
}

func NewTaskHandler(ts TaskService, dh taskdelivery.DeliveryHandler) Handler {
	return Handler{
		DeliveryHandler: dh,
		TaskService:     ts,
	}
}

func (h Handler) HandleTask(ctx context.Context, newEvent event.ProcessedEvent) error {
	var task taskentity.Task
	var err error

	var taskStatus taskentity.IntegrationDeliveryStatus

	taskID := newEvent.ID()

	// Get task status using idempotency in the task service.
	if taskStatus, err = h.TaskService.GetTaskStatusByID(ctx, taskID); err != nil {
		// todo use richError
		return err
	}

	if !taskStatus.CanBeExecuted() {
		slog.Debug(fmt.Sprintf("Task [%s] has %s status and is not executable", taskID, taskStatus.String()))

		return nil
	}

	if taskStatus.IsBroadcast() {
		// Get task info such as attempts, failed reason and... using repository.
		task, err = h.TaskService.GetTaskByID(taskID)
		if err != nil {
			return err
		}
	} else {
		task = taskentity.MakeTaskUsingProcessedEvent(newEvent)
	}

	// DeliveryHandler is responsible for delivering event to third party destinations.
	// DeliveryHandler should consider max_retries base on integration configs.
	deliveryResponse, err := h.DeliveryHandler.Handle(task, newEvent)
	if err != nil {
		return err
	}

	// DeliveryStatus is set in delivery handler in case of success, retriable failed, unretriable failed.
	task.IntegrationDeliveryStatus = deliveryResponse.DeliveryStatus
	// Attempts is incremented in delivery handler in case of success.
	task.Attempts = deliveryResponse.Attempts
	// FailedReason describes what caused the failure.
	task.FailedReason = deliveryResponse.FailedReason

	err = h.TaskService.UpsertTaskAndSaveIdempotency(ctx, task)
	if err != nil {
		// todo what should we do if error occurs in updating task repo or idempotency ?
		return err
	}

	return nil
}
