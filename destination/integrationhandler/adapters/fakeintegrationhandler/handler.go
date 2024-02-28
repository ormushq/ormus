package fakeintegrationhandler

import (
	"github.com/ormushq/ormus/event"
	"log"
	"time"
)

type FakeHandler struct {
}

func New() *FakeHandler {
	return &FakeHandler{}
}

func (h FakeHandler) Handle(e event.ProcessedEvent) error {
	time.Sleep(2 * time.Second)

	log.Printf("\033[32mIntegration Message [%s::%s] handled successfully.!\033[0m\n", e.MessageID, e.Integration.ID)

	return nil
}
