package channel

type Adapter interface {
	GetInputChannel(name string) (chan<- []byte, error)
	GetOutputChannel(name string) (<-chan Message, error)
	GetMode(name string) (Mode, error)
	NewChannel(name string, mode Mode, bufferSize, numberInstants, maxRetryPolicy int) error
}

type Message struct {
	Ack  func() error
	Body []byte
}
