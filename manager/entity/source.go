package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// TODO: need change feilds
type Source struct {
	WriteKey    ulid.ULID
	Name        string
	Description string
	ProjectId   string
	OwnerId     string
	Status      bool
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}
