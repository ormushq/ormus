package entity

import (
	"time"
)

type WriteKey string // because we might change the format in future

// TODO: need change feilds.
type Source struct {
	SourceID    string
	WriteKey    WriteKey
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      bool
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}
