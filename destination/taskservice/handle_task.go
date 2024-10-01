package taskservice

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (s Service) HandleTask(ctx context.Context, newEvent event.ProcessedEvent) error {
	// TODO: complete retruns
	var task taskentity.Task
	var err error
	tracer := otela.NewTracer("taskservice")
	ctx, span := tracer.Start(ctx, "taskservice@HandleTask")
	defer span.End()
	newEvent.TracerCarrier = otela.GetCarrierFromContext(ctx)
	span.AddEvent("start-handle-task")

	taskID := newEvent.ID()

	unlock, err := s.LockTaskByID(ctx, taskID)
	if err != nil {
		span.AddEvent("error-on-lock-task", trace.WithAttributes(
			attribute.String("error", err.Error())))

		return err
	}
	defer func() {
		unlockErr := unlock()
		if unlockErr != nil {
			logger.L().Error(fmt.Sprintf("unlock task failed %s", unlockErr))
		}
	}()

	var taskStatus taskentity.IntegrationDeliveryStatus

	// Get task status using idempotency in the task service.
	if taskStatus, err = s.GetTaskStatusByID(ctx, taskID); err != nil {
		span.AddEvent("error-on-get-task-status", trace.WithAttributes(
			attribute.String("error", err.Error())))

		// todo use richError
		return err
	}

	// If task status is not executable, we don't need to do anything.
	if !taskStatus.CanBeExecuted() {
		slog.Debug(fmt.Sprintf("Task [%s] has %s status and is not executable", taskID, taskStatus.String()))
		span.AddEvent("task-can`t-be-executed")

		return nil
	}

	span.AddEvent("task-status-retrieved", trace.WithAttributes(
		attribute.String("status", taskStatus.String()),
	))

	span.AddEvent("task-ready-to-execute")

	// If task status is broadcast, we need to get task info from repository.
	if taskStatus.IsBroadcast() {
		span.AddEvent("task-status-is-broadcast-try-to-get-task-from-repo")
		task, err = s.GetTaskByIDWithContext(ctx, taskID)
		task.ProcessedEvent = newEvent
		if err != nil {
			span.AddEvent("get-task-by-id-error", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))

			return err
		}
	} else {
		span.AddEvent("task-status-is-not-broadcast-try-to-create-task-from-event")
		// If task status is not broadcast, we need to create task using processed event.
		task = taskentity.MakeTaskUsingProcessedEvent(newEvent)
	}

	// DeliveryHandler is responsible for delivering processed event to third party destinations.
	destinationType := task.DestinationSlug()

	span.AddEvent("destination-type", trace.WithAttributes(
		attribute.String("destinationType", string(destinationType)),
	))

	// Get Delivery handler from task delivery mapper.
	deliveryHandler, ok := taskdelivery.Mapper[destinationType]
	if !ok {
		span.AddEvent("there-is-no-delivery-handler-for-this-type")

		return fmt.Errorf("destination type %s is not supported", destinationType)
	}

	// Deliver processed event to third party destinations using corresponding delivery handler.
	deliveryResponse, err := deliveryHandler.Handle(task)
	if err != nil {
		span.AddEvent("delivery-handler-error", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))

		return err
	}

	// DeliveryStatus is set in delivery handler in case of success, retriable failed, unretriable failed.
	task.DeliveryStatus = deliveryResponse.DeliveryStatus
	// Attempts is incremented in delivery handler in case of success.
	task.Attempts = deliveryResponse.Attempts
	// FailedReason describes what caused the failure.
	task.FailedReason = deliveryResponse.FailedReason

	failedReason := ""
	if deliveryResponse.FailedReason != nil {
		failedReason = *deliveryResponse.FailedReason
	}
	span.AddEvent("delivery-handler-response", trace.WithAttributes(
		attribute.String("DeliveryStatus", deliveryResponse.DeliveryStatus.String()),
		attribute.String("Attempts", strconv.Itoa(int(deliveryResponse.Attempts))),
		attribute.String("FailedReason", failedReason),
	))
	err = s.UpsertTaskAndSaveIdempotency(ctx, task)
	// dispatch event of success delivery
	if err != nil {
		// todo what should we do if error occurs in updating task repo or idempotency ?
		return err
	}

	return nil
}
