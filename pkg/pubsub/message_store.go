package pubsub

import "sync"

// just storge message in this map for guaranteed delivery of messages.
type InMemoryMessageStore struct {
	Data map[string][]publisherMessage
	Mu   sync.RWMutex
}

type publisherMessage struct {
	m *Message
	p *Publisher
}
