package worker

import (
	"github.com/ormushq/ormus/destination/handler"
	"github.com/ormushq/ormus/destination/queue"
	"time"
)

// Worker represents a worker that processes jobs from the queue.
type Worker struct {
	ID                 int
	Queue              queue.JobQueue
	IntegrationHandler handler.IntegrationHandler
}

func New(id int, q queue.JobQueue, h handler.IntegrationHandler) *Worker {
	return &Worker{
		ID:                 id,
		IntegrationHandler: h,
		Queue:              q,
	}
}

// ProcessJobs continuously processes jobs from the queue.
func (w Worker) ProcessJobs() {
	for {
		job := w.Queue.Dequeue(w.IntegrationHandler.GetTopic())
		if job.ID == 0 {
			// No job received, sleep for a while before checking again
			time.Sleep(1 * time.Second)
			continue
		}

		// todo handle job using topic handler
		//todo workers should handle idempotency
		w.IntegrationHandler.Handle(&job)

	}
}
