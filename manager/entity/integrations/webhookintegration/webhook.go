package webhookintegration

type WebhookMethod string

const (
	POSTWebhookMethod  WebhookMethod = "POST"
	PUTWebhookMethod   WebhookMethod = "PUT"
	PATCHWebhookMethod WebhookMethod = "PATCH"
)

type WebhookConfig struct {
	Headers map[string]string
	Payload map[string]string
	Method  WebhookMethod
	URL     string
}
