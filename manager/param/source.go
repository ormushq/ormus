package param

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type AddSourceRequest struct {
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
}

type AddSourceResponse struct {
	WriteKey    ulid.ULID
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      bool
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}

type UpdateSourceRequest struct {
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      bool
}

type UpdateSourceResponse struct {
	WriteKey    ulid.ULID
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
	Status      bool
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}
