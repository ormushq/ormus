package taskservice

import (
	"context"
	"github.com/ormushq/ormus/adapter/otela"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskservice/param"
)

type Repository interface {
	GetTaskByIDWithContext(ctx context.Context, taskID string) (taskentity.Task, error)
	UpsertTaskWithContext(ctx context.Context, taskID string, request param.UpsertTaskRequest) error
}

type Idempotency interface {
	GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error)
	SaveTaskStatus(ctx context.Context, taskID string, status taskentity.IntegrationDeliveryStatus) error
}

type Locker interface {
	Lock(ctx context.Context, key string, ttl int64) (unlock func() error, err error)
}

type Service struct {
	idempotency Idempotency
	repo        Repository
	locker      Locker
}

func New(idempotency Idempotency, repo Repository, l Locker) Service {
	return Service{
		idempotency: idempotency,
		repo:        repo,
		locker:      l,
	}
}

func (s Service) LockTaskByID(ctx context.Context, taskID string) (unlock func() error, err error) {
	tracer := otela.NewTracer("taskservice")
	ctx, span := tracer.Start(ctx, "taskservice@LockTaskByID", trace.WithAttributes(
		attribute.String("taskId", taskID)))
	defer span.End()

	lockKey := "task:" + taskID
	const ttl = 10

	return s.locker.Lock(ctx, lockKey, ttl)
}

func (s Service) GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	tracer := otela.NewTracer("taskservice")
	ctx, span := tracer.Start(ctx, "taskservice@GetTaskStatusByID", trace.WithAttributes(
		attribute.String("taskId", taskID)))
	defer span.End()

	span.AddEvent("start-get-task-status")
	// Acquire a lock with a 10-second TTL
	// todo get lock prefix and ttl from config

	status, err := s.idempotency.GetTaskStatusByID(ctx, taskID)
	if err != nil {
		span.AddEvent("error-on-get-status", trace.WithAttributes(
			attribute.String("error", err.Error())))

		return taskentity.InvalidTaskStatus, err
	}
	span.AddEvent("status-retrieved", trace.WithAttributes(
		attribute.String("status", status.String())))

	return status, nil
}

func (s Service) GetTaskByID(taskID string) (taskentity.Task, error) {
	return s.GetTaskByIDWithContext(context.Background(), taskID)
}

func (s Service) GetTaskByIDWithContext(ctx context.Context, taskID string) (taskentity.Task, error) {
	tracer := otela.NewTracer("taskservice")
	_, span := tracer.Start(ctx, "taskservice@GetTaskByIDWithContext", trace.WithAttributes(
		attribute.String("taskId", taskID)))
	defer span.End()

	task, err := s.repo.GetTaskByIDWithContext(ctx, taskID)
	if err != nil {
		return taskentity.Task{}, err
	}
	span.AddEvent("task-retrieved")

	return task, nil
}

func (s Service) UpsertTaskAndSaveIdempotency(ctx context.Context, t taskentity.Task) error {
	tracer := otela.NewTracer("taskservice")
	ctx, span := tracer.Start(ctx, "taskservice@UpsertTaskAndSaveIdempotency")
	defer span.End()

	req := param.UpsertTaskRequest{
		IntegrationDeliveryStatus: t.DeliveryStatus,
		Attempts:                  t.Attempts,
		FailedReason:              t.FailedReason,
		UpdatedAt:                 time.Now(),
	}

	rErr := s.repo.UpsertTaskWithContext(ctx, t.ID, req)
	if rErr != nil {
		return rErr
	}
	span.AddEvent("task-upserted")

	iErr := s.idempotency.SaveTaskStatus(ctx, t.ID, t.DeliveryStatus)
	if iErr != nil {
		// todo is it better to rollback updated task status in idempotency?
		return iErr
	}
	span.AddEvent("task-status-updated")

	return nil
}
