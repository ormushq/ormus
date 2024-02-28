package processedevent

type Consumer interface {
	Consume() error
	Close() error
}
