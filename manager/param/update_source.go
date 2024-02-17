package param

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
)

type UpdateSourceRequest struct {
	Name        string
	Description string
	ProjectID   string
	Status      entity.Status
}

type UpdateSourceResponse struct {
	ID          string
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
