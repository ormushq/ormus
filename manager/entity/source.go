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
	ID          string         `json:"id"`
	TokenID     string         `json:"token_id"`
	WriteKey    WriteKey       `json:"write_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ProjectID   string         `json:"project_id"`
	OwnerID     string         `json:"owner_id"`
	Status      Status         `json:"status"`
	Metadata    SourceMetadata `json:"metadata"`
	Deleted     bool           `json:"deleted"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at"`
}

type SourceMetadata struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Slug     string         `json:"slug"`
	Category SourceCategory `json:"category"`
}
