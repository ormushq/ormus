package params

import (
	"github.com/ormushq/ormus/event"
	"time"
)

type TrackEventRequest struct {
	MessageID event.ID
	Type      event.Type // track, page , ...

	// Category  string // specific to page event
	Name string

	Properties *event.Properties

	Integration *event.Integrations
	Ctx         *event.Context

	SendAt            time.Time
	ReceivedAt        time.Time
	OriginalTimeStamp time.Time // TODO: don't know the difference between this and Timestamp example(2023-11-19T12:04:25.271Z)
	Timestamp         time.Time // 2023-11-19T12:04:25.779Z

	AnonymousID event.ID
	UserID      event.ID
	GroupID     event.ID
	PreviousID  event.ID

	Event string

	WriteKey string
	MetaData event.MetaData // TODO: all fields had ambiguity so did not added yet

	Options *event.Options
}

type TrackEventResponse struct {
	Event event.CoreEvent
}
