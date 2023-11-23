package event

import "time"

type Type string

const (
	TrackEventType    Type = "track"
	PageEventType     Type = "page"
	IdentityEventType Type = "identity"
	GroupEventType    Type = "group"
	AliasEventType    Type = "alias"
	ScreenEventType   Type = "screen"
)

type CoreEvent struct {
	MessageID ID
	Type      Type // track, page , ...

	// Category  string // specific to page event
	Name string

	Properties *Properties

	Integration *Integrations
	Ctx         *Context

	SendAt            time.Time
	ReceivedAt        time.Time
	OriginalTimeStamp time.Time // TODO: don't know the difference between this and Timestamp example(2023-11-19T12:04:25.271Z)
	Timestamp         time.Time // 2023-11-19T12:04:25.779Z

	AnonymousID ID
	UserID      ID
	GroupID     ID
	PreviousID  ID

	Event string

	WriteKey string
	MetaData MetaData // TODO: all fields had ambiguity so did not added yet

	Options *Options
}

type MetaData struct {
}

type Options struct {
	Integrations *Integrations
	Timestamp    time.Time
	Ctx          *Context
	AnonymousID  ID
	UserID       ID
	Traits       Traits
	CustomData   *CustomData
}

type Integrations struct {
	All    bool
	Config Properties // TODO: i don't think this is a proper name for this field but, didn't have a better name on my mind
}
