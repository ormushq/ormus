package fakeintegrationhandler

import (
	"log"
	"time"

	"github.com/ormushq/ormus/event"
)

const sleepTime = 2

type FakeHandler struct{}

func New() *FakeHandler {
	return &FakeHandler{}
}

func (h FakeHandler) Handle(e event.ProcessedEvent) error {
	time.Sleep(sleepTime * time.Second)

	log.Printf("\033[32mIntegration Message [%s::%s] handled successfully.!\033[0m\n", e.MessageID, e.Integration.ID)

	return nil
}
