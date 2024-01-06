package param

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
)

type ResponseIntegration struct {
	Name           string
	Category       entity.Category
	Status         bool
	Source         entity.Source
	Type           string
	ConnectionType entity.ConnectionType
	CreatedAt      time.Time
}