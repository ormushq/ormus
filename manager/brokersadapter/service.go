package brokersadapter

// MessageBroker defines the interface for a message broker client.
type MessageBroker interface {
	// PublishMessage publishes messages to a specified topic.
	PublishMessage(topic string, messages ...*Message) error

	// ConsumeMessage consumes messages on the specified topic.
	ConsumeMessage(topic string) (<-chan *Message, error)

	// Close closes the message broker client.
	Close() error
}
