package pubsub

// just storge message in this map for guaranteed delivery of messages.
type InMemoryMessageStore struct {
	message         *Message
	pubslisher      *Publisher
	subscriberCount int
}
