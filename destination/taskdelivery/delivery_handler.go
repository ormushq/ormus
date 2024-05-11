package taskdelivery

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
)

// DeliveryHandler is responsible for delivering processed event to third party destinations.
type DeliveryHandler interface {
	Handle(task taskentity.Task) (param.DeliveryTaskResponse, error)
}
