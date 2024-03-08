package entity

import (
	"time"
)

type DestinationCategory string

type DestinationType string

const (
	WebhookDestinationType DestinationType = "webhook"
)

const (
	Analytics      DestinationCategory = "analytics"
	Advertising    DestinationCategory = "advertising"
	CRM            DestinationCategory = "crm"
	EmailMarketing DestinationCategory = "email-marketing"
	Livechat       DestinationCategory = "livechat"
	Payments       DestinationCategory = "payments"
	Surveys        DestinationCategory = "Surveys"
)

// ConnectionType each third party destination are compatible with one of these methods
// it means we have to deliver data to the destinations with these methods
// https://github.com/ormushq/ormus/issues/9
type ConnectionType string

const (
	EventStream ConnectionType = "event-stream"
	Storage     ConnectionType = "storage"
	ReversETL   ConnectionType = "reverse-ETL"
)

// Integration is a connector that allows our app send data to an external service or application.
type Integration struct {
	ID             string
	SourceID       string
	Name           string
	Metadata       DestinationMetadata
	ConnectionType ConnectionType
	Enabled        bool
	Config         Config
	CreatedAt      time.Time
}

type Config interface{}

type DestinationMetadata struct {
	ID         string
	Name       string          // webhook, Google Universal Analytics
	Slug       DestinationType // webhook, google-analytics
	Categories []DestinationCategory
}
