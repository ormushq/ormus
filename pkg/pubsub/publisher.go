package pubsub

type Publisher struct {
	name string
}

func NewPublisher(name string) *Publisher {
	return &Publisher{
		name: name,
	}
}
