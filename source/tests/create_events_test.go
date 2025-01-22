package tests

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/encoder"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvents(t *testing.T) {
	c := GetConfigs()
	wg := sync.WaitGroup{}
	ctx := context.Background()
	events := []event.CoreEvent{
		{
			WriteKey:   faker.NAME,
			Type:       faker.NAME,
			Name:       "event1",
			SendAt:     time.Now(),
			ReceivedAt: time.Now(),
			Event:      faker.NAME,
			Timestamp:  time.Now(),
			Properties: (*event.Properties)(&map[string]string{
				"user": "ali",
			}),
		},
		{
			WriteKey:   faker.ID,
			Type:       faker.NAME,
			Name:       "event2",
			SendAt:     time.Now(),
			ReceivedAt: time.Now(),
			Event:      faker.NAME,
			Timestamp:  time.Now(),
			Properties: (*event.Properties)(&map[string]string{
				"user": "ali",
			}),
		},
	}

	_, err := c.EventRepo.CreateNewEvent(ctx, events, &wg, c.Cfg.Source.NewEventQueueName)
	if err != nil {
		logger.L().Error(err.Error())
	}
	result := []struct {
		Count int `db:"count"`
	}{}

	q := c.SylladbConn.Query("SELECT count(*) FROM test_source.event;", []string{})
	err = q.Select(&result)
	if err != nil {
		logger.L().Error(err.Error())
	}
	assert.Equal(t, len(events), result[0].Count)
	ch, _ := c.brokerAdapter.GetOutputChannel(c.Cfg.Source.NewEventQueueName)
	for range 2 {
		tmp := <-ch
		decodedEvent := encoder.DecodeNewEvent(string(tmp.Body))
		assert.Contains(t, []string{events[0].Name, events[1].Name}, decodedEvent.Name)
		assert.Contains(t, []string{events[0].WriteKey, events[1].WriteKey}, decodedEvent.WriteKey)
	}
}
