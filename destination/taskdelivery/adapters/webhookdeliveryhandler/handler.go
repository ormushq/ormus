package webhookdeliveryhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/manager/entity"
	"net/http"
	"net/url"
)

type WebhookHandler struct{}

func New() *WebhookHandler {
	return &WebhookHandler{}
}

func (h WebhookHandler) Handle(e event.ProcessedEvent) error {
	var response *http.Response
	var err error

	// TODO: The methods have a lot of duplicated code and need to be cleaned up a bit
	switch e.Integration.Config.Method {
	case entity.GETWebhookMethod:
		response, err = WebhookGetHandler(e.Integration.Config)
		if err != nil {
			fmt.Println("WebhookGetHandler error : ", err)
		}
	case entity.POSTWebhookMethod:
		response, err = WebhookPostHandler(e.Integration.Config)
		if err != nil {
			fmt.Println("WebhookGetHandler error : ", err)
		}
	}

	fmt.Println(response.Status)
	return nil
}

func WebhookPostHandler(config entity.WebhookConfig) (*http.Response, error) {
	payloadMap := make(map[string]string)
	for _, item := range config.Payload {
		payloadMap[item.Key] = item.Value
	}

	payloadJSON, err := json.Marshal(payloadMap)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(string(config.Method), config.Url, bytes.NewBuffer(payloadJSON))
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

func WebhookGetHandler(config entity.WebhookConfig) (*http.Response, error) {
	client := &http.Client{}

	u, err := url.Parse(config.Url)
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
