package taskentity

import (
	"encoding/json"
	"time"

	"github.com/ormushq/ormus/contract/goprotobuf/processedevent"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"google.golang.org/protobuf/proto"
)

type IntegrationDeliveryStatus uint8

const (
	InvalidTaskStatus IntegrationDeliveryStatus = iota
	NotExecutedTaskStatus
	RetriableFailedTaskStatus
	UnRetriableFailedTaskStatus
	SuccessTaskStatus
)

// Task represents a delivering processed event to corresponding third party integrations.
type Task struct {
	ID             string
	DeliveryStatus IntegrationDeliveryStatus
	Attempts       uint8
	FailedReason   *string
	ProcessedEvent event.ProcessedEvent
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (t IntegrationDeliveryStatus) String() string {
	return []string{"invalid", "not_executed", "retriable_failed", "unretriable_failed", "success"}[t]
}

func (t IntegrationDeliveryStatus) ToNumericString() string {
	return []string{"0", "1", "2", "3", "4"}[t]
}

func (t IntegrationDeliveryStatus) IsValid() bool {
	return t > InvalidTaskStatus && t <= SuccessTaskStatus
}

func (t IntegrationDeliveryStatus) CanBeExecuted() bool {
	return t == NotExecutedTaskStatus || t == RetriableFailedTaskStatus
}

func (t IntegrationDeliveryStatus) IsBroadcast() bool {
	return t != InvalidTaskStatus && t != NotExecutedTaskStatus
}

func MakeTaskUsingProcessedEvent(pe event.ProcessedEvent) Task {
	return Task{
		ID:             pe.ID(),
		ProcessedEvent: pe,
		DeliveryStatus: NotExecutedTaskStatus,
		Attempts:       0,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}

func (t Task) DestinationSlug() entity.DestinationType {
	return t.ProcessedEvent.Integration.Metadata.Slug
}

func UnmarshalBytesToTask(body []byte) (Task, error) {
	var t Task
	logger.L().Debug(string(body))
	if err := json.Unmarshal(body, &t); err != nil {
		logger.L().Error("Error on unMarshaling task :", err)

		return Task{}, err
	}

	return t, nil
}

func UnmarshalBytesToProcessedEvent(body []byte) (event.ProcessedEvent, error) {
	var pe event.ProcessedEvent
	if err := json.Unmarshal(body, &pe); err != nil {
		logger.L().Error("Error on unMarshaling processed event :", err)

		return event.ProcessedEvent{}, err
	}

	return pe, nil
}

func ProtoUnmarshalBytesToProcessedEvnet(body []byte) (*processedevent.ProcessedEvent, error) {
	var pe processedevent.ProcessedEvent

	if err := proto.Unmarshal(body, &pe); err != nil {
		logger.L().Error("Error on unMarshaling processed event :", err)

		return &pe, err
	}

	return &pe, nil
}

func NumericStringToIntegrationDeliveryStatus(statusStr string) IntegrationDeliveryStatus {
	switch statusStr {
	case "0":
		return InvalidTaskStatus
	case "1":
		return NotExecutedTaskStatus
	case "2":
		return RetriableFailedTaskStatus
	case "3":
		return UnRetriableFailedTaskStatus
	case "4":
		return SuccessTaskStatus
	default:
		return InvalidTaskStatus
	}
}
