package taskentity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ormushq/ormus/event"
)

type IntegrationDeliveryStatus uint8

const (
	InvalidTaskStatus IntegrationDeliveryStatus = iota
	NotExecutedTaskStatus
	RetriableFailedTaskStatus
	UnRetriableFailedTaskStatus
	Success
)

// Task represents a delivering processed event to corresponding third party integrations.
type Task struct {
	ID                        string
	IntegrationDeliveryStatus IntegrationDeliveryStatus
	Attempts                  uint8
	FailedReason              *string
	ProcessedEvent            event.ProcessedEvent
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func (t IntegrationDeliveryStatus) String() string {
	return []string{"invalid", "not_executed", "retriable_failed", "unretriable_failed", "success"}[t]
}

func (t IntegrationDeliveryStatus) IsValid() bool {
	return t > InvalidTaskStatus && t <= Success
}

func (t IntegrationDeliveryStatus) CanBeExecuted() bool {
	return t == NotExecutedTaskStatus || t == RetriableFailedTaskStatus
}

func (t IntegrationDeliveryStatus) IsBroadcast() bool {
	return t != InvalidTaskStatus && t != NotExecutedTaskStatus
}

func MakeTaskUsingProcessedEvent(pe event.ProcessedEvent) *Task {
	return &Task{
		ID:                        pe.MessageID + "-" + pe.Integration.ID,
		ProcessedEvent:            pe,
		IntegrationDeliveryStatus: NotExecutedTaskStatus,
		Attempts:                  1,
		CreatedAt:                 time.Time{},
		UpdatedAt:                 time.Time{},
	}
}

func (t Task) DestinationSlug() string {
	return t.ProcessedEvent.Integration.Metadata.Slug
}

func UnmarshalBytesToProcessedEvent(body []byte) (event.ProcessedEvent, error) {
	var pe event.ProcessedEvent
	if err := json.Unmarshal(body, &pe); err != nil {
		fmt.Println("Error on unMarshaling processed event :", err)

		return event.ProcessedEvent{}, err
	}

	return pe, nil
}
