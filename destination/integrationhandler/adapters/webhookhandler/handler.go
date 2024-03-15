package webhookhandler

import (
	"log"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
	"github.com/ormushq/ormus/event"
)

type WebhookHandler struct{}

func (h WebhookHandler) Handle(e event.ProcessedEvent) (param.HandleTaskResponse, error) {
	log.Printf("Message ID : %s", e.MessageID)
	// todo fill job.ProcessedEvent.Integration.Config using redis/cache and grpc (consider fallback approach)
	// todo Handle webhook integration (Send http request)
	res := param.HandleTaskResponse{
		ErrorReason:    nil,
		DeliveryStatus: taskentity.Success,
	}

	return res, nil
}
