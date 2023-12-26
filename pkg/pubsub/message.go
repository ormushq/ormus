package pubsub

import "github.com/google/uuid"

type Message struct {
	ID        uuid.UUID
	Payload   any
	Publisher Publisher
}

func NewMessage(id uuid.UUID, payload any) *Message {
	return &Message{
		ID:      id,
		Payload: payload,
	}
}
