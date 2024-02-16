package main

import (
	"github.com/ormushq/ormus/destination/adapter/inmemoryqueue"
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/destination/handler/adapters/webhook/fakewebhookhandler"
	"github.com/ormushq/ormus/destination/queue"
	"github.com/ormushq/ormus/destination/queue/worker"
	"sync"
	"time"
)

func main() {

	//-----Setup queue and workers-----

	// Create an in-memory job queue
	inMemoryQueue := &inmemoryqueue.Queue{}

	// Create workers for different topics
	workers := []*worker.Worker{
		worker.New(1, inMemoryQueue, fakewebhookhandler.WebhookHandler{}),
		worker.New(2, inMemoryQueue, fakewebhookhandler.WebhookHandler{}),
	}

	// Start workers
	wg := sync.WaitGroup{}
	for _, wrk := range workers {
		wg.Add(1)
		go func(w *worker.Worker) {
			defer wg.Done()
			w.ProcessJobs()
		}(wrk)
	}

	//----- Consume processed events -----

	// Fake publish and consumer
	for i := 1; i <= 10; i++ {
		//generate fake processedEvent
		processedEvent := entity.ProcessedEvent{}
		// Enqueue webhook jobs
		webhookJob := queue.Job{
			ID:             i,
			Topic:          "webhook",
			ProcessedEvent: processedEvent,
			Timestamp:      time.Now(),
		}
		inMemoryQueue.Enqueue(webhookJob)
	}

	// Wait for workers to finish
	wg.Wait()
}
