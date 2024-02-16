package fakewebhookhandler

import (
	"fmt"
	"github.com/ormushq/ormus/destination/queue"
)

// WebhookHandler is a handler for webhook jobs.
type WebhookHandler struct {
}

// Handle handles the webhook job.
func (h WebhookHandler) Handle(job *queue.Job) {
	fmt.Printf("Processing webhook job %d", job.ID)
	// todo fill job.ProcessedEvent.Integration.Config using IntegrationClient
	// todo Implement logic for processing webhook job.
}

func (h WebhookHandler) GetTopic() string {
	return "webhook"
}
