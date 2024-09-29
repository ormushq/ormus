package entity

import (
	"time"
)

type WriteKey string // because we might change the format in future

type SourceCategory string

type Status string

const (
	SourceStatusActive    Status = "active"
	SourceStatusNotActive Status = "not active"
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
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type SourceMetadata struct {
	ID       string
	Name     string
	Slug     string
	Category SourceCategory
}
