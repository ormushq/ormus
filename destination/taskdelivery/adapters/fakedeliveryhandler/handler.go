package fakedeliveryhandler

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
	"github.com/ormushq/ormus/event"
)

type FakeHandler struct{}

func New() *FakeHandler {
	return &FakeHandler{}
}

const fakeProcessingTimeSecond = 2

func (h FakeHandler) Handle(t taskentity.Task, _ event.ProcessedEvent) (param.DeliveryTaskResponse, error) {
	time.Sleep(fakeProcessingTimeSecond * time.Second)

	slog.Info(fmt.Sprintf("Task [%s] handled successfully!", t.ID))

	res := param.DeliveryTaskResponse{
		Attempts:       1,
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
	}

	return res, nil
}
