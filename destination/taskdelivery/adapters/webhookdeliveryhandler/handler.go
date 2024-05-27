package webhookdeliveryhandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/taskdelivery/param"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	"github.com/ormushq/ormus/pkg/richerror"
)

type WebhookHandler struct{}

func New() *WebhookHandler {
	return &WebhookHandler{}
}

func (h WebhookHandler) Handle(task taskentity.Task) (param.DeliveryTaskResponse, error) {
	const op = "webhookhandler.Handle"

	_ = fmt.Sprintf("I am here in %s ------------------------ \n", op)

	// TODO: use oneof
	config, ok := task.ProcessedEvent.Integration.Config.(webhookintegration.WebhookConfig)
	if !ok {
		logger.L().Info("invalid configuration for webhook")

		return param.DeliveryTaskResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).
			WithMessage("invalid configuration for webhook")
	}

	_, err := MakeHTTPRequest(config)
	if err != nil {
		logger.L().Error("error in webhookhandler.Handle when try to Do GET request", err)

		return param.DeliveryTaskResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage("unexpected error when try to do GET webhook request")
	}

	return param.DeliveryTaskResponse{
		FailedReason:   nil,
		Attempts:       0,
		DeliveryStatus: taskentity.SuccessTaskStatus,
	}, nil
}

func MakeHTTPRequest(config webhookintegration.WebhookConfig) (*http.Response, error) {
	const op = "webhookdeliveryhandler.MakeHTTPRequest"

	payloadMap := make(map[string]string)
	for k, v := range config.Payload {
		payloadMap[k] = v
	}

	payloadJSON, err := json.Marshal(payloadMap)
	if err != nil {
		return nil, err
	}

	// TODO: read timeout from config
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, string(config.Method), config.URL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, err
	}

	// TODO: check headers in segment again
	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.L().Info(fmt.Sprintf("failed to close http response body in %s", op))
		}
	}(response.Body)

	return response, nil
}
