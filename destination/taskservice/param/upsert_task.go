package param

import (
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
)

type UpsertTaskRequest struct {
	IntegrationDeliveryStatus taskentity.IntegrationDeliveryStatus
	Attempts                  uint8
	FailedReason              *string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}
