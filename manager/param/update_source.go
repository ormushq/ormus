package param

import (
	"time"
)

type UpdateSourceRequest struct {
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      bool
}

type UpdateSourceResponse struct {
	SourceID    string
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
