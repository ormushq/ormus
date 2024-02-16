package entity

type DestinationType string

type IntegrationConfig map[string]interface{}

const (
	Webhook DestinationType = "webhook"
)

// Integration is a connector that allows our app send data to an external service or application.
type Integration struct {
	ID            string
	SourceID      string
	DestinationID string
	Enabled       bool
	Config        IntegrationConfig
}
