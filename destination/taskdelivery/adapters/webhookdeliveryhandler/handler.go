package webhookdeliveryhandler

import (
	"bytes"
	"encoding/json"
	"github.com/ormushq/ormus/destination/entity/taskentity"
	"github.com/ormushq/ormus/destination/integrationhandler/param"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	"github.com/ormushq/ormus/pkg/richerror"
	"net/http"
)

type WebhookHandler struct{}

func New() *WebhookHandler {
	return &WebhookHandler{}
}

// Handle TODO: why we have processedEvent here?! since we have processedEvent in task
// Handle TODO: why we should have error and fail reason on HandleTaskResponse
func (h WebhookHandler) Handle(task taskentity.Task, processedEvent event.ProcessedEvent) (param.HandleTaskResponse, error) {
	const op = "webhookhandler.Handle"

	config, ok := task.ProcessedEvent.Integration.Config.(webhookintegration.WebhookConfig)
	if !ok {
		logger.L().Info("invalid configuration for webhook")

		return param.HandleTaskResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).
			WhitMessage("invalid configuration for webhook")
	}

	_, err := MakeHTTPRequest(config)
	if err != nil {
		logger.L().Error("error in webhookhandler.Handle when try to Do GET request", err)

		return param.HandleTaskResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected).
			WhitMessage("unexpected error when try to do GET webhook request")
	}

	return param.HandleTaskResponse{
		FailedReason:   nil,
		Attempts:       0,
		DeliveryStatus: taskentity.SuccessTaskStatus,
	}, nil
}

func MakeHTTPRequest(config webhookintegration.WebhookConfig) (*http.Response, error) {
	payloadMap := make(map[string]string)
	for _, item := range config.Payload {
		payloadMap[item.Key] = item.Value
	}

	payloadJSON, err := json.Marshal(payloadMap)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(string(config.Method), config.URL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, err
	}

	for _, header := range config.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, nil
}
