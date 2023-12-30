package param

import "time"

type AddSourceRequest struct {
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
}

type AddSourceResponse struct {
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
