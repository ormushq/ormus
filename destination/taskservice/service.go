package taskservice

import (
	"context"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskservice/param"
)

type Repository interface {
	GetTaskByID(taskID string) (taskentity.Task, error)
	UpsertTask(taskID string, request param.UpsertTaskRequest) error
}

type Idempotency interface {
	GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error)
	SaveTaskStatus(ctx context.Context, taskID string, status taskentity.IntegrationDeliveryStatus) error
}

type Service struct {
	idempotency Idempotency
	repo        Repository
}

func New(idempotency Idempotency, repo Repository) Service {
	return Service{
		idempotency: idempotency,
		repo:        repo,
	}
}

func (s Service) GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	status, err := s.idempotency.GetTaskStatusByID(ctx, taskID)
	if err != nil {
		return taskentity.InvalidTaskStatus, err
	}

	return status, nil
}

func (s Service) GetTaskByID(taskID string) (taskentity.Task, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return taskentity.Task{}, err
	}

	return task, nil
}

func (s Service) UpsertTaskAndSaveIdempotency(ctx context.Context, t taskentity.Task) error {
	req := param.UpsertTaskRequest{
		IntegrationDeliveryStatus: t.IntegrationDeliveryStatus,
		Attempts:                  t.Attempts,
		FailedReason:              t.FailedReason,
		UpdatedAt:                 time.Now(),
	}

	rErr := s.repo.UpsertTask(t.ID, req)
	if rErr != nil {
		return rErr
	}

	iErr := s.idempotency.SaveTaskStatus(ctx, t.ID, t.IntegrationDeliveryStatus)
	if iErr != nil {
		// todo is it better to rollback updated task status in idempotency?
		return iErr
	}

	return nil
}
