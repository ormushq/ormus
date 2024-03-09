package taskidempotencyservice

import (
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/event"
)

type ServiceImpl struct {
	repo Repository
}

type Repository interface {
	UpsertTask(taskID string, status entity.TaskStatus) error
	GetStatusByTaskID(taskID string) (entity.TaskStatus, error)
}

type Service interface {
	GetTaskID(pe *event.ProcessedEvent) string
	Save(taskID string, status entity.TaskStatus) error
	GetStatusByTaskID(taskID string) (entity.TaskStatus, error)
	IntegrationHandlerIsEnable(taskID string) (bool, error)
}

func New(repo Repository) ServiceImpl {
	return ServiceImpl{repo: repo}
}

func (s ServiceImpl) GetTaskID(pe *event.ProcessedEvent) string {
	return pe.MessageID + "::" + pe.Integration.ID
}

func (s ServiceImpl) Save(taskID string, status entity.TaskStatus) error {
	return s.repo.UpsertTask(taskID, status)
}

func (s ServiceImpl) GetStatusByTaskID(taskID string) (entity.TaskStatus, error) {
	return s.repo.GetStatusByTaskID(taskID)
}

func (s ServiceImpl) IntegrationHandlerIsEnable(taskID string) (bool, error) {
	status, err := s.GetStatusByTaskID(taskID)
	if err != nil {
		return false, err
	}

	// A task doesn't exist in the idempotency system or has failed in previous integration handlers.
	if status == entity.NotExists || status == entity.FailedInIntegrationHandler {
		return true, nil
	}

	return false, nil
}
