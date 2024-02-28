package entity

import "time"

type IntegrationConfig map[string]any

type DestinationCategory string

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
	Config         IntegrationConfig
	CreatedAt      time.Time
}

type DestinationMetadata struct {
	ID         string
	Name       string // Javascript, Google Universal Analytics
	Slug       string // javascript, google-analytics
	Categories []DestinationCategory
}
