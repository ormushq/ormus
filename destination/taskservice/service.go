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

func (s Service) GetTaskStatusByID(ctx context.Context, taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	// Acquire a lock with a 10-second TTL
	// todo get lock prefix and ttl from config
	lockKey := "task:" + taskID
	const ttl = 10
	unlock, err := s.locker.Lock(ctx, lockKey, ttl)
	if err != nil {
		return taskentity.InvalidTaskStatus, err
	}

	defer func() {
		err = unlock()
	}()

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
		IntegrationDeliveryStatus: t.DeliveryStatus,
		Attempts:                  t.Attempts,
		FailedReason:              t.FailedReason,
		UpdatedAt:                 time.Now(),
	}

	rErr := s.repo.UpsertTask(t.ID, req)
	if rErr != nil {
		return rErr
	}

	iErr := s.idempotency.SaveTaskStatus(ctx, t.ID, t.DeliveryStatus)
	if iErr != nil {
		// todo is it better to rollback updated task status in idempotency?
		return iErr
	}

	return nil
}
