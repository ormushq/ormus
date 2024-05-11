package webhookdeliveryhandler

import (
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
)

type WebhookHandler struct{}

func (h WebhookHandler) Handle(t taskentity.Task) (param.DeliveryTaskResponse, error) {
	// todo webhook handler is responsible for sending processed event to url of webhook and making DeliveryTaskResponse
	// todo webhook handler should consider max_retry_exceeded and other necessary policies in e.Integration.Config

	res := param.DeliveryTaskResponse{
		FailedReason:   nil,
		DeliveryStatus: taskentity.SuccessTaskStatus,
		Attempts:       t.Attempts + 1, // in every send request we should increment attempts value
	}

	return res, nil
}
