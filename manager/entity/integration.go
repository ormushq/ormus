package entity

import (
	"time"
)

type Integration struct {
	Name             string
	Category         Category
	Status           bool
	Source           Source
	Type             string
	ConnectionType   ConnectionType
	LatestSyncStatus time.Time
	CreatedAt        time.Time
}

// Category The integrations we have are categorized into groups
// Facebook Pixel and Google Ads (Classic) are placed in the Advertising category
type Category string

const (
	Analytics      Category = "analytics"
	Advertising    Category = "advertising"
	CRM            Category = "crm"
	EmailMarketing Category = "email-marketing"
	Livechat       Category = "livechat"
	Payments       Category = "payments"
	Surveys        Category = "Surveys"
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
