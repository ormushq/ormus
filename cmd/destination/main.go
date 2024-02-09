package main

import (
	"github.com/go-faker/faker/v4"
	"github.com/ormushq/ormus/destination/broker"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
)

func generateFakeProcessedEvent() broker.ProcessedEvent {
	var e event.CoreEvent
	var i entity.Integration

	faker.FakeData(&e)
	faker.FakeData(&i)

	return broker.ProcessedEvent{
		Event:       e,
		Integration: []entity.Integration{i},
	}
}

func main() {
	// TODO: get entranceEvent from pub/sub
	entranceEvent := generateFakeProcessedEvent()
}
