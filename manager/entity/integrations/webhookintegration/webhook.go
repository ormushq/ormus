package webhookintegration

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

	// what about time out retryable
	// general config
	PaxBroadcast uint
	MaxAttempt   uint
}
