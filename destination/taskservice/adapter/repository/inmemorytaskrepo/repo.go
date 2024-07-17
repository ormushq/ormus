package inmemorytaskrepo

import (
	"context"
	"fmt"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskservice/param"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (db DB) GetTaskByID(taskID string) (taskentity.Task, error) {
	return db.GetTaskByIDWithContext(context.Background(), taskID)
}

func (db DB) GetTaskByIDWithContext(ctx context.Context, taskID string) (taskentity.Task, error) {
	tracer := otela.NewTracer("inmemorytaskrepo")
	_, span := tracer.Start(ctx, "inmemorytaskrepo@GetTaskByIDWithContext", trace.WithAttributes(
		attribute.String("taskID", taskID),
	))
	defer span.End()

	task, ok := db.tasks[taskID]
	if !ok {
		span.AddEvent("task-not-found")

		return taskentity.Task{}, fmt.Errorf("task not found")
	}
	span.AddEvent("task-fetched")

	return *task, nil
}

func (db DB) UpsertTask(taskID string, request param.UpsertTaskRequest) error {
	return db.UpsertTaskWithContext(context.Background(), taskID, request)
}

func (db DB) UpsertTaskWithContext(ctx context.Context, taskID string, request param.UpsertTaskRequest) error {
	tracer := otela.NewTracer("inmemorytaskrepo")
	_, span := tracer.Start(ctx, "inmemorytaskrepo@UpsertTaskWithContext", trace.WithAttributes(
		attribute.String("taskID", taskID),
	))
	defer span.End()

	db.mutex.Lock()
	span.AddEvent("db-tasks-locked")
	defer db.mutex.Unlock()
	task, ok := db.tasks[taskID]
	if !ok {
		span.AddEvent("task-not-exist-try-to-create")
		// todo create task
		task = &taskentity.Task{
			ID: taskID,
		}
		db.tasks[taskID] = task
		span.AddEvent("task-created")
	}

	task.DeliveryStatus = request.IntegrationDeliveryStatus
	task.Attempts = request.Attempts
	task.FailedReason = request.FailedReason

	return nil
}
