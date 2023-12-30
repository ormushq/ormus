package entity

import "time"
import "github.com/ormushq/ormus/manager/entity"

type Integration struct {
	Name             string
	CategoryID       Category
	Status           bool
	Source           entity.Source
	Type             string
	ConnectionType   ConnectionType
	CreatedAt        *time.Time
	LatestSyncStatus *time.Time
}
