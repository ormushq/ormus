package event

import (
	"github.com/ormushq/ormus/manager/entity"
	"time"
)

type ProcessedEvent struct {
	//required fields
	SourceID    string // Write key or sourceID shows that which source, events coming from
	Integration entity.Integration
	MessageID   string // Automatically collected by Segment, a unique identifier for each message
	EventType   Type   // Type of message, corresponding to the API method: 'identify', 'group', 'track', 'page', 'screen' or 'alias'.
	Version     uint8  // The version of the Tracking API

	SentAt            time.Time // Time on client device when call was sent or sentAt value manually passed in.
	ReceivedAt        time.Time // Time on server clock when call was received
	OriginalTimestamp time.Time // Time on the client device when call was invoked
	Timestamp         time.Time // Calculated by Server to correct client-device clock skew using the following formula: receivedAt - (sentAt - originalTimestamp)

	//optional fields
	UserID      *string     // Unique identifier for the user in your database. A userId or an anonymousId is required.
	AnonymousId *string     // A uniqueID substitute for a User ID, for cases when you don’t have an absolutely unique identifier.
	Event       *string     // track
	Name        *string     // page | screen
	GroupID     *string     // group
	PreviousId  *string     // alias
	Context     *Context    // Dictionary of extra information that provides useful context about a message
	Properties  *Properties // Custom information about the event
	Traits      *Traits     // identify | group
}
