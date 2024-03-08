package webhookdeliveryhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity/integrations/webhookintegration"
	"github.com/ormushq/ormus/pkg/richerror"
	"net/http"
	"net/url"
)

type WebhookHandler struct{}

func New() *WebhookHandler {
	return &WebhookHandler{}
}

func (h WebhookHandler) Handle(e event.ProcessedEvent) error {
	const op = "webhookhandler.Handle"

	var response *http.Response
	var err error

	config, ok := e.Integration.Config.(webhookintegration.WebhookConfig)
	if !ok {
		logger.L().Info("invalid configuration for webhook")

		return richerror.New(op).WithKind(richerror.KindInvalid).
			WhitMessage("invalid configuration for webhook")
	}

	// TODO: The methods have a lot of duplicated code and need to be cleaned up a bit
	switch config.Method {
	case webhookintegration.GETWebhookMethod:
		response, err = WebhookGetHandler(config)
		if err != nil {
			logger.L().Error("error in webhookhandler.Handle when try to Do GET request", err)

			return richerror.New(op).WithKind(richerror.KindUnexpected).
				WhitMessage("unexpected error when try to do GET webhook request")
		}
	case webhookintegration.POSTWebhookMethod:
		response, err = WebhookPostHandler(config)
		if err != nil {
			logger.L().Error("error in webhookhandler.Handle when try to Do POST request", err)

			return richerror.New(op).WithKind(richerror.KindUnexpected).
				WhitMessage("unexpected error when try to do POST webhook request")
		}
	}

	fmt.Println(response.StatusCode)
	return nil
}

func WebhookPostHandler(config webhookintegration.WebhookConfig) (*http.Response, error) {
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

func WebhookGetHandler(config webhookintegration.WebhookConfig) (*http.Response, error) {
	client := &http.Client{}

	u, err := url.Parse(config.URL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for _, item := range config.Payload {
		q.Add(item.Key, item.Value)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(string(config.Method), u.String(), nil)
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
