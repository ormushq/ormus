package observer

import "github.com/google/uuid"

type Message struct {
	Id      uuid.UUID
	Payload any
}

func NewMessage(id uuid.UUID, payload any) *Message {

	return &Message{
		Id:      id,
		Payload: payload,
	}
}
