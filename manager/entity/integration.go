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

type ConnectionType string

const (
	EventStream ConnectionType = "event-stream"
	Storage     ConnectionType = "storage"
	ReversETL   ConnectionType = "reverse-ETL"
)
