package taskservice

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery"
	"github.com/ormushq/ormus/event"
)

func (s Service) HandleTask(ctx context.Context, newEvent event.ProcessedEvent) error {
	var task taskentity.Task
	var err error

	var taskStatus taskentity.IntegrationDeliveryStatus
	taskID := newEvent.ID()

	// Get task status using idempotency in the task service.
	if taskStatus, err = s.GetTaskStatusByID(ctx, taskID); err != nil {
		// todo use richError
		return err
	}

	// If task status is not executable, we don't need to do anything.
	if !taskStatus.CanBeExecuted() {
		slog.Debug(fmt.Sprintf("Task [%s] has %s status and is not executable", taskID, taskStatus.String()))

		return nil
	}

	// If task status is broadcast, we need to get task info from repository.
	if taskStatus.IsBroadcast() {
		task, err = s.GetTaskByID(taskID)
		task.ProcessedEvent = newEvent
		if err != nil {
			return err
		}
	} else {
		// If task status is not broadcast, we need to create task using processed event.
		task = taskentity.MakeTaskUsingProcessedEvent(newEvent)
	}

	// DeliveryHandler is responsible for delivering processed event to third party destinations.
	destinationType := task.DestinationSlug()

	// Get Delivery handler from task delivery mapper.
	deliveryHandler, ok := taskdelivery.Mapper[destinationType]

	fmt.Printf("%v", destinationType)

	if !ok {
		return fmt.Errorf("destination type %s is not supported", destinationType)
	}

	// Deliver processed event to third party destinations using corresponding delivery handler.
	deliveryResponse, err := deliveryHandler.Handle(task)
	if err != nil {
		return err
	}

	// DeliveryStatus is set in delivery handler in case of success, retriable failed, unretriable failed.
	task.DeliveryStatus = deliveryResponse.DeliveryStatus
	// Attempts is incremented in delivery handler in case of success.
	task.Attempts = deliveryResponse.Attempts
	// FailedReason describes what caused the failure.
	task.FailedReason = deliveryResponse.FailedReason

	err = s.UpsertTaskAndSaveIdempotency(ctx, task)
	// dispatch event of success delivery
	if err != nil {
		// todo what should we do if error occurs in updating task repo or idempotency ?
		return err
	}

	return nil
}
