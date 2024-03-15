package fakeintegrationhandler

import (
	"log"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
)

type FakeHandler struct{}

func New() *FakeHandler {
	return &FakeHandler{}
}

const fakeProcessingTimeSecond = 2

func (h FakeHandler) Handle(t *taskentity.Task) (param.HandleTaskResponse, error) {
	time.Sleep(fakeProcessingTimeSecond * time.Second)

	// TODO consider max_retry_exceeded, is_broadcasted and other necessary configs

	log.Printf("\033[32mTask [%s] handled successfully.!\033[0m\n", t.ID)

	res := param.HandleTaskResponse{
		ErrorReason:    nil,
		DeliveryStatus: taskentity.Success,
	}

	return res, nil
}
