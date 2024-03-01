package entity

type Header struct {
	Key   string
	Value string
}

type WebhookMethod string

const (
	POSTWebhookMethod WebhookMethod = "POST"
	GETWebhookMethod  WebhookMethod = "GET"
)

type Payload struct {
	Key   string
	Value string
}

type WebhookConfig struct {
	Headers []Header
	Payload []Payload
	Method  WebhookMethod
	Url     string
}
