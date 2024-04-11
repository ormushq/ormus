package webhookintegration

type WebhookMethod string

const (
	POSTWebhookMethod  WebhookMethod = "POST"
	PUTWebhookMethod   WebhookMethod = "PUT"
	PATCHWebhookMethod WebhookMethod = "PATCH"
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
