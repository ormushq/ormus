package webhookhandler

import (
	"log"

	"github.com/ormushq/ormus/event"
)

type WebhookHandler struct{}

func (h WebhookHandler) Handle(e event.ProcessedEvent) error {
	log.Printf("Message ID : %s", e.MessageID)
	// todo fill job.ProcessedEvent.Integration.Config using redis/cache and grpc (consider fallback approach)
	// todo Handle webhook integration (Send http request)

	return nil
}
