package entity

import (
	"time"
)

type (
	ID         string
	EventType  string
	Properties map[string]interface{}
)

const (
	TrackEventType    EventType = "track"    // Track API call is how you record any actions your users perform, along with any properties that describe the action.
	PageEventType     EventType = "page"     // Page call lets you record whenever a user sees a page of your website, along with any optional properties about the page
	IdentityEventType EventType = "identity" // Identity lets you tie a user to their actions and record traits about them
	GroupEventType    EventType = "group"    // Group API call is how you associate an individual user with a group, such as a company, organization, account, project, or team.
	AliasEventType    EventType = "alias"    // Alias method is an advanced method used to merge 2 unassociated user identities, effectively connecting 2 sets of user data in one profile.
	ScreenEventType   EventType = "screen"   // Screen call lets you record whenever a user sees a screen, the mobile equivalent of Page, in your mobile app, along with any properties about the screen.
)

type ProcessedEvent struct {
	//required fields
	WriteKey    string // Write key or sourceID shows that which source, events coming from
	Integration Integration
	MessageID   string    // Automatically collected by Segment, a unique identifier for each message
	EventType   EventType // Type of message, corresponding to the API method: 'identify', 'group', 'track', 'page', 'screen' or 'alias'.
	Version     uint8     // The version of the Tracking API

	SentAt            time.Time // Time on client device when call was sent or sentAt value manually passed in.
	ReceivedAt        time.Time // Time on server clock when call was received
	OriginalTimestamp time.Time // Time on the client device when call was invoked
	Timestamp         time.Time // Calculated by Server to correct client-device clock skew using the following formula: receivedAt - (sentAt - originalTimestamp)

	//optional fields
	UserID      string                 // Unique identifier for the user in your database. A userId or an anonymousId is required.
	AnonymousId string                 // A pseudo-unique substitute for a User ID, for cases when you don’t have an absolutely unique identifier.
	Event       string                 // track
	Name        string                 // page | screen
	GroupID     string                 // group
	PreviousId  string                 // alias
	Context     map[string]interface{} // Dictionary of extra information that provides useful context about a message
	Properties  Properties             // Custom information about the event
	Traits      map[string]interface{} // identify | group

}
