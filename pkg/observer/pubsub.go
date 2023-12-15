package observer

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

type Config struct {
	SubscriberChannelBufferSize int64
}

var Pubsub *GoChannelPubsub

func init() {}

type GoChannelPubsub struct {
	config Config

	subscribersWg          sync.WaitGroup
	subscribers            map[string][]*Subscriber
	subscribersLock        sync.RWMutex
	subscribersByTopicLock sync.Map // map of *sync.Mutex

	closed     bool
	closedLock sync.Mutex
	closing    chan struct{}
}

func NewPubsub(config Config) *GoChannelPubsub {
	return &GoChannelPubsub{
		config:                 config,
		subscribers:            make(map[string][]*Subscriber),
		subscribersByTopicLock: sync.Map{},
		closing:                make(chan struct{}),
	}
}

func (g *GoChannelPubsub) Publish(topic string, messages ...*Message) error {
	if g.isClosed() {
		return fmt.Errorf("pubsub closed")
	}

	g.subscribersLock.RLock()
	defer g.subscribersLock.RUnlock()

	topicLock, _ := g.subscribersByTopicLock.LoadOrStore(topic, &sync.Mutex{})
	tl, ok := topicLock.(*sync.Mutex)
	if !ok {
		panic("Type assertion failed in pubsub.publish")
	}
	tl.Lock()
	defer tl.Unlock()

	for i := range messages {

		msg := messages[i]

		err := g.sendMessage(topic, msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GoChannelPubsub) sendMessage(topic string, message *Message) error {
	subscribers := g.topicSubscribers(topic)

	if len(subscribers) == 0 {
		return fmt.Errorf("no subscribers have")
	}

	go func(subscribers []*Subscriber) {
		wg := &sync.WaitGroup{}

		for i := range subscribers {
			subscriber := subscribers[i]

			wg.Add(1)
			go func() {
				subscriber.sendMessageToSubscriber(message)
				wg.Done()
			}()
		}

		wg.Wait()
	}(subscribers)

	return nil
}

func (g *GoChannelPubsub) topicSubscribers(topic string) []*Subscriber {
	subscribers, ok := g.subscribers[topic]
	if !ok {
		return nil
	}

	// let's do a copy to avoid race conditions and deadlocks due to lock
	subscribersCopy := make([]*Subscriber, len(subscribers))
	copy(subscribersCopy, subscribers)

	return subscribersCopy
}

func (g *GoChannelPubsub) Subscribe(topic string) (chan *Message, error) {
	g.closedLock.Lock()

	if g.closed {
		g.closedLock.Unlock()

		return nil, fmt.Errorf("Pub/Sub closed")
	}

	g.subscribersWg.Add(1)
	g.closedLock.Unlock()

	g.subscribersLock.Lock()

	topicLock, _ := g.subscribersByTopicLock.LoadOrStore(topic, &sync.Mutex{})
	tl, ok := topicLock.(*sync.Mutex)
	if !ok {
		panic("Type assertion failed in pubsub.publish")
	}
	tl.Lock()

	s := &Subscriber{
		id:             uuid.New(),
		messageChannel: make(chan *Message, g.config.SubscriberChannelBufferSize),
		closing:        make(chan struct{}),
	}

	defer g.subscribersLock.Unlock()
	defer tl.Unlock()

	g.addSubscriber(topic, s)

	return s.messageChannel, nil
}

func (g *GoChannelPubsub) addSubscriber(topic string, s *Subscriber) {
	if _, ok := g.subscribers[topic]; !ok {
		g.subscribers[topic] = make([]*Subscriber, 0)
	}

	g.subscribers[topic] = append(g.subscribers[topic], s)
}

func (g *GoChannelPubsub) isClosed() bool {
	g.closedLock.Lock()
	defer g.closedLock.Unlock()

	return g.closed
}

// Close closes the GoChannel Pub/Sub.
func (g *GoChannelPubsub) Close() error {
	g.closedLock.Lock()
	defer g.closedLock.Unlock()

	if g.closed {
		return nil
	}

	g.closed = true
	close(g.closing)

	// log
	g.subscribersWg.Wait()

	return nil
}

type Subscriber struct {
	id uuid.UUID

	sending        sync.Mutex
	messageChannel chan *Message

	closed  bool
	closing chan struct{}
}

func (s *Subscriber) Close() {
	if s.closed {
		return
	}

	close(s.closing)

	// ensuring that we are not sending to closed channel
	s.sending.Lock()
	defer s.sending.Unlock()

	s.closed = true

	close(s.messageChannel)
}

func (s *Subscriber) sendMessageToSubscriber(msg *Message) {
	s.sending.Lock()
	defer s.sending.Unlock()

	if s.closed {
		slog.Info("Pub/Sub closed, discarding msg ")

		return
	}

	select {

	case s.messageChannel <- msg:
		slog.Info("Sent message to subscriber ")

	case <-s.closing:
		slog.Info("Closing, message discarded ")

		return
	}
}
