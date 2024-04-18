package webhookdeliveryhandler

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
	"github.com/ormushq/ormus/event"
)

type WebhookHandler struct{}

func (h WebhookHandler) Handle(t taskentity.Task, pe event.ProcessedEvent) (param.DeliveryTaskResponse, error) {
	// todo webhook handler is responsible for sending processed event to url of webhook and making DeliveryTaskResponse
	// todo webhook handler should consider max_retry_exceeded and other necessary policies in e.Integration.Config

	// get configs from processed event
	println(pe.Integration.Config)

	res := param.DeliveryTaskResponse{
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
		Attempts:       t.Attempts + 1, // in every send request we should increment attempts value
	}

	return res, nil
}
