package integrationhandler

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
	"github.com/ormushq/ormus/event"
)

// IntegrationHandler defines the interface for a topic handler.
type IntegrationHandler interface {
	Handle(task taskentity.Task, processedEvent event.ProcessedEvent) (param.HandleTaskResponse, error)
}
