package integrationhandler

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
)

// IntegrationHandler defines the interface for a topic handler.
type IntegrationHandler interface {
	Handle(task *taskentity.Task) (param.HandleTaskResponse, error)
}
