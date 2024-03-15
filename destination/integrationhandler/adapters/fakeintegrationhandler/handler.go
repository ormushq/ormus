package fakeintegrationhandler

import (
	"log"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
	"github.com/ormushq/ormus/event"
)

type FakeHandler struct{}

func New() *FakeHandler {
	return &FakeHandler{}
}

const fakeProcessingTimeSecond = 2

func (h FakeHandler) Handle(t taskentity.Task, _ event.ProcessedEvent) (param.HandleTaskResponse, error) {
	time.Sleep(fakeProcessingTimeSecond * time.Second)

	log.Printf("\033[32mTask [%s] handled successfully.!\033[0m\n", t.ID)

	res := param.HandleTaskResponse{
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
	}

	return res, nil
}
