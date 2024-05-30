package mock

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/richerror"
	"time"
)

const RepoErr = "repository error"

type DefaultEventTest struct {
	MessageID         event.ID
	Type              event.Type
	Name              string
	Properties        *event.Properties
	Integration       *event.Integrations
	Ctx               *event.Context
	SendAt            time.Time
	ReceivedAt        time.Time
	OriginalTimeStamp time.Time
	Timestamp         time.Time
	AnonymousID       event.ID
	UserID            event.ID
	GroupID           event.ID
	PreviousID        event.ID
	Event             string
	WriteKey          string
	MetaData          event.MetaData
	Options           *event.Options
}

func DefaultEvent() DefaultEventTest {
	return DefaultEventTest{
		MessageID:         "1",
		Type:              "track",
		Name:              "track",
		Properties:        nil,
		Integration:       nil,
		Ctx:               nil,
		SendAt:            time.Time{},
		ReceivedAt:        time.Time{},
		OriginalTimeStamp: time.Time{},
		Timestamp:         time.Time{},
		AnonymousID:       "1",
		UserID:            "1",
		GroupID:           "1",
		PreviousID:        "1",
		Event:             "track",
		WriteKey:          "123456789",
		MetaData:          event.MetaData{},
		Options:           nil,
	}
}

type MockRepo struct {
	events []event.CoreEvent
	hasErr bool
}

func NewMockRepository(hasErr bool) MockRepo {
	var events []event.CoreEvent
	defaultEvent := DefaultEvent()
	events = append(events,
		event.CoreEvent{
			MessageID:         defaultEvent.MessageID,
			Type:              defaultEvent.Type,
			Name:              defaultEvent.Name,
			Properties:        nil,
			Integration:       nil,
			Ctx:               nil,
			SendAt:            time.Time{},
			ReceivedAt:        time.Time{},
			OriginalTimeStamp: time.Time{},
			Timestamp:         time.Time{},
			AnonymousID:       defaultEvent.AnonymousID,
			UserID:            defaultEvent.UserID,
			GroupID:           defaultEvent.GroupID,
			PreviousID:        defaultEvent.PreviousID,
			Event:             defaultEvent.Event,
			WriteKey:          defaultEvent.WriteKey,
			MetaData:          event.MetaData{},
			Options:           nil,
		})

	return MockRepo{
		events: events,
		hasErr: hasErr,
	}
}
func (m *MockRepo) InsertEvent(ctx context.Context, e event.CoreEvent) (event.CoreEvent, error) {
	if m.hasErr {
		return event.CoreEvent{}, richerror.New("MockRepo.InsertEvent").WithWrappedError(fmt.Errorf(RepoErr))
	}

	m.events = append(m.events, e)

	return e, nil
}
