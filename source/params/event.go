package params

import (
	"time"
)

type TrackEventRequest struct {
	Type       string                 `json:"type" validate:"required"`
	Name       string                 `json:"name" validate:"required,min=3,max=255"`
	SendAt     time.Time              `json:"send_at" validate:"required"`
	ReceivedAt time.Time              `json:"received_at" validate:"required"`
	Timestamp  time.Time              `json:"timestamp" validate:"required"`
	Event      string                 `json:"event" validate:"required"`
	WriteKey   string                 `json:"write_key" validate:"required"`
	Properties map[string]interface{} `json:"properties" validate:"required"`
}

type TrackEventResponse struct {
	ID string `json:"id"`
}
