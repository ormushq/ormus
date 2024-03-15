package webhookhandler

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
	"github.com/ormushq/ormus/event"
)

type WebhookHandler struct{}

func (h WebhookHandler) Handle(t taskentity.Task, pe event.ProcessedEvent) (param.HandleTaskResponse, error) {
	// todo webhook handler is responsible for sending processed event to url of webhook and making HandleTaskResponse
	// todo webhook handler should consider max_retry_exceeded and other necessary policies in e.Integration.Config

	// get configs from processed event
	println(pe.Integration.Config)

	res := param.HandleTaskResponse{
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
		Attempts:       t.Attempts + 1, // in every send request we should increment attempts value
	}

	return res, nil
}
