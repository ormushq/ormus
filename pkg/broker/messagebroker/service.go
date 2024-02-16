package messagebroker

// Consumer defines the interface for a message consumer.
type Consumer interface {
	// ConsumeMessage consumes messages on the specified topic.
	ConsumeMessage(topic string) (<-chan *Message, error)

	// Close closes the consumer.
	Close() error
}

// Publisher defines the interface for a message publisher.
type Publisher interface {
	// PublishMessage publishes messages to a specified topic.
	PublishMessage(topic string, exchangeName string, messages ...*Message) error

	// Close closes the publisher.
	Close() error
}

// BrokerService represents a service that interacts with a message broker.
type BrokerService struct {
	Consumer  Consumer  // Represents the consumer interface for interacting with the message broker.
	Publisher Publisher // Represents the publisher interface for interacting with the message broker.
}

// NewBroker creates a new instance of BrokerService with the provided consumer and publisher implementations.
func NewBroker(consumer Consumer, publisher Publisher) *BrokerService {
	return &BrokerService{
		Consumer:  consumer,
		Publisher: publisher,
	}
}
