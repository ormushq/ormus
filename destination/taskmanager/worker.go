package taskmanager

import (
	"sync"
)

// Worker represents a worker that processes jobs of queue.
type Worker interface {
	Run(done <-chan bool, wg *sync.WaitGroup)
}
