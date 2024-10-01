package protobufmapper

import (
	"github.com/ormushq/ormus/contract/go/task"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MapTaskFromProtobuf TODO: implement me
func MapTaskFromProtobuf(t *task.Task) taskentity.Task {
	return taskentity.Task{}
}

// MapTaskToProtobuf TODO: implement me
func MapTaskToProtobuf(t taskentity.Task) *task.Task {
	theFailedReason := ""
	if t.FailedReason != nil {
		theFailedReason = *t.FailedReason
	}

	return &task.Task{
		Id:             t.ID,
		TaskStatus:     task.TaskStatus(t.DeliveryStatus),
		Attempts:       uint32(t.Attempts),
		FailedReason:   theFailedReason,
		ProcessedEvent: MapProcessedEventToProtobuf(t.ProcessedEvent),
		CreatedAt:      timestamppb.New(t.CreatedAt),
		UpdatedAt:      timestamppb.New(t.UpdatedAt),
	}
}
