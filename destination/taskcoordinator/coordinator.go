package taskcoordinator

import (
	"sync"

	"github.com/ormushq/ormus/event"
)

type Coordinator interface {
	Start(processedEvents <-chan event.ProcessedEvent, done <-chan bool, wg *sync.WaitGroup)
}
