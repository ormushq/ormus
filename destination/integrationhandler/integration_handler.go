package integrationhandler

import (
	"github.com/ormushq/ormus/event"
)

// IntegrationHandler defines the interface for a topic handler.
type IntegrationHandler interface {
	Handle(processedEvent event.ProcessedEvent) error
}
