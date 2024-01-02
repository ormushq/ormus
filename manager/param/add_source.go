package param

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
)

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
	Status      entity.Status
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    *time.Time
}
