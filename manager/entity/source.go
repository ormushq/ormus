package entity

import (
	"time"
)

type WriteKey string // because we might change the format in future

<<<<<<< HEAD
type SourceCategory string

=======
>>>>>>> fb1e3b5 (feat(manager): add new source (#46))
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
<<<<<<< HEAD
	Metadata    SourceMetadata
=======
>>>>>>> fb1e3b5 (feat(manager): add new source (#46))
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}
<<<<<<< HEAD

type SourceMetadata struct {
	ID       string
	Name     string
	Slug     string
	Category SourceCategory
}
=======
>>>>>>> fb1e3b5 (feat(manager): add new source (#46))
