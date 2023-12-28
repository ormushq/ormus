package entity

import (
	"time"
)

// TODO: need change feilds.
type Source struct {
	WriteKey    string
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      bool
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}
