package taskdelivery

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
	"github.com/ormushq/ormus/event"
)

// DeliveryHandler is responsible for delivering event to third party destinations.
type DeliveryHandler interface {
	Handle(task taskentity.Task, processedEvent event.ProcessedEvent) (param.DeliveryTaskResponse, error)
}
