package param

import "github.com/ormushq/ormus/destination/entity/taskentity"

type HandleTaskResponse struct {
	ErrorReason    *string
	DeliveryStatus taskentity.IntegrationDeliveryStatus
}
