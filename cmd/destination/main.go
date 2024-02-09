package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
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
	// TODO: clean code and use setupServices function to set up all configurations
	Redis, err := redis.New(config.C().Redis)
	if err != nil {
		panic(fmt.Sprintf("We have a problem in the cache db: %v", err))
	}

	// TODO: get entranceEvent from pub/sub
	entranceEvent := generateFakeProcessedEvent()

	for _, integration := range entranceEvent.Integration {
		err := eventOrchestration(entranceEvent.Event, integration)
		if err != nil {
			// TODO: log error
		}
	}
}

func eventOrchestration(event event.CoreEvent, integrations entity.Integration) error {
	//TODO; check cache for event id and integration id if status is failed or success than update layer 11 with event id
	// else decide to send event with integration to the right queue
	return nil
}
