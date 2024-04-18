package param

import "github.com/ormushq/ormus/destination/entity/taskentity"

type DeliveryTaskResponse struct {
	FailedReason   *string
	Attempts       uint8
	DeliveryStatus taskentity.IntegrationDeliveryStatus
}
