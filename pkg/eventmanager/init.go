package eventmanager

import (
	"context"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/contract/go/internalevent"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func NewWriteKeyGeneratedEvent(ctx context.Context, ev *internalevent.WriteKeyGeneratedEvent) *internalevent.Event {
	if ctx == nil {
		ctx = context.WithoutCancel(context.Background())
	}

	return &internalevent.Event{
		EventName:     internalevent.EventName_EVENT_NAME_WRITE_KEY_GENERATED,
		Time:          timestamppb.New(time.Now()),
		TracerCarrier: otela.GetCarrierFromContext(ctx),
		Payload:       &internalevent.Event_WriteKeyGeneratedEvent{WriteKeyGeneratedEvent: ev},
	}
}

func NewUserCreatedEvent(ctx context.Context, ev *internalevent.UserCreatedEvent) *internalevent.Event {
	if ctx == nil {
		ctx = context.WithoutCancel(context.Background())
	}

	return &internalevent.Event{
		EventName:     internalevent.EventName_EVENT_NAME_USER_CREATED,
		Time:          timestamppb.New(time.Now()),
		TracerCarrier: otela.GetCarrierFromContext(ctx),
		Payload:       &internalevent.Event_UserCreatedEvent{UserCreatedEvent: ev},
	}
}

func NewProjectCreatedEvent(ctx context.Context, ev *internalevent.ProjectCreatedEvent) *internalevent.Event {
	if ctx == nil {
		ctx = context.WithoutCancel(context.Background())
	}

	return &internalevent.Event{
		EventName:     internalevent.EventName_EVENT_NAME_PROJECT_CREATED,
		Time:          timestamppb.New(time.Now()),
		TracerCarrier: otela.GetCarrierFromContext(ctx),
		Payload:       &internalevent.Event_ProjectCreatedEvent{ProjectCreatedEvent: ev},
	}
}

func NewTaskCreatedEvent(ctx context.Context, ev *internalevent.TaskCreatedEvent) *internalevent.Event {
	if ctx == nil {
		ctx = context.WithoutCancel(context.Background())
	}

	return &internalevent.Event{
		EventName:     internalevent.EventName_EVENT_NAME_TASK_CREATED,
		Time:          timestamppb.New(time.Now()),
		TracerCarrier: otela.GetCarrierFromContext(ctx),
		Payload:       &internalevent.Event_TaskCreatedEvent{TaskCreatedEvent: ev},
	}
}
