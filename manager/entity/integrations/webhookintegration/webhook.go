package webhookintegration

type WebhookMethod string

const (
	POSTWebhookMethod WebhookMethod = "POST"
	GETWebhookMethod  WebhookMethod = "GET"
)

type Header struct {
	Key   string
	Value string
}

type Payload struct {
	Key   string
	Value string
}

type WebhookConfig struct {
	Headers []Header
	Payload []Payload
	Method  WebhookMethod
	URL     string
}
