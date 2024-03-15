package taskrouter

import (
	"sync"

	"github.com/ormushq/ormus/event"
)

// Coordinator is responsible for setup task managers and publish coming process events using suitable task publishers.
type Coordinator interface {
	Start(processedEvents <-chan *event.ProcessedEvent, done <-chan bool, wg *sync.WaitGroup)
}
