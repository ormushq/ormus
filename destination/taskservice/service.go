package taskservice

import (
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskservice/param"
)

type Repository interface {
	GetTaskByID(taskID string) (*taskentity.Task, error)
	UpsertTask(taskID string, request param.UpsertTaskRequest) error
}

type Idempotency interface {
	GetTaskStatusByID(taskID string) (taskentity.IntegrationDeliveryStatus, error)
	SaveTaskStatus(taskID string, status taskentity.IntegrationDeliveryStatus) error
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

func (s Service) GetTaskStatusByID(taskID string) (taskentity.IntegrationDeliveryStatus, error) {
	status, err := s.idempotency.GetTaskStatusByID(taskID)
	if err != nil {
		return taskentity.InvalidTaskStatus, err
	}

	return status, nil
}

func (s Service) GetTaskByID(taskID string) (*taskentity.Task, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s Service) UpsertTaskAndSaveIdempotency(t taskentity.Task, newStatus taskentity.IntegrationDeliveryStatus, failedReason *string) error {
	req := param.UpsertTaskRequest{
		IntegrationDeliveryStatus: newStatus,
		Attempts:                  t.Attempts + 1,
		FailedReason:              failedReason,
		UpdatedAt:                 time.Now(),
	}

	rErr := s.repo.UpsertTask(t.ID, req)
	if rErr != nil {
		return rErr
	}

	iErr := s.idempotency.SaveTaskStatus(t.ID, newStatus)
	if iErr != nil {
		// todo I think it is better to rollback updated task in this scenario.
		return iErr
	}

	return nil
}
