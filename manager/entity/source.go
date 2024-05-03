package entity

import (
	"time"
)

type WriteKey string // because we might change the format in future

type SourceCategory string

type Status string

const (
	StatusActive    Status = "active"
	StatusNotActive Status = "not active"
)

// TODO: need change feilds.
type Source struct {
	ID          string
	WriteKey    WriteKey
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      Status
	Metadata    SourceMetadata
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}

type SourceMetadata struct {
	ID       string
	Name     string
	Slug     string
	Category SourceCategory
}
