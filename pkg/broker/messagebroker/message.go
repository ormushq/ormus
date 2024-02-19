package messagebroker

import "github.com/google/uuid"

// Message represents a message in the pub/sub system.
type Message struct {
	ID      uuid.UUID
	Topic   string
	Payload []byte
}

// NewMessage creates a new message with the given topic, payload, and metadata.
func NewMessage(topic string, payload []byte) Message {
	return Message{
		ID:      uuid.New(),
		Topic:   topic,
		Payload: payload,
	}
}
