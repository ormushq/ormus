package pubsub

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/logger"
)

type Config struct {
	SubscriberChannelBufferSize int64 `koanf:"subscriber_channel_buffer_size"`
}

var GoPubSub *PubSub

func StartPubSub(cfg Config) {
	GoPubSub = NewPubSub(cfg)
}

type PubSub struct {
	cfg Config

	subscribersWg          sync.WaitGroup
	subscribers            map[string][]*subscriber
	subscribersLock        sync.RWMutex
	subscribersByTopicLock sync.Map // map of *sync.Mutex

	closed     bool
	closedLock sync.Mutex

	messageRepo InMemoryMessageStore
}

func NewPubSub(cfg Config) *PubSub {
	return &PubSub{
		cfg:                    cfg,
		subscribers:            make(map[string][]*subscriber),
		subscribersByTopicLock: sync.Map{},
		messageRepo:            InMemoryMessageStore{Data: make(map[string][]publisherMessage)},
	}
}

// publsh messages in topic.
func (g *PubSub) Publish(publisher *Publisher, topic string, messages ...*Message) error {
	// check pubsub is closed
	if g.isClosed() {
		logger.L().Info("pubsub closed")
		panic("pubsub closed")
	}

	// lock subscribersLock because publisher cannot send messages at the same time
	g.subscribersLock.RLock()
	defer g.subscribersLock.RUnlock()

	// lock topic because publisher can not send message at the same time
	topicLock, _ := g.subscribersByTopicLock.LoadOrStore(topic, &sync.Mutex{})
	tl, ok := topicLock.(*sync.Mutex)
	if !ok {
		logger.L().Info("type assertion failed in pubsub.publish")
		panic("type assertion failed in pubsub.publish")
	}

	tl.Lock()
	defer tl.Unlock()

	unsendedMessages, ok := g.messageRepo.Data[topic]
	// if there are messages ok == true
	if ok {
		for _, msg := range unsendedMessages {
			// then we need check did this publisher send this message or not?
			if msg.p == publisher {
				// append to new messages list to send again
				messages = append(messages, msg.m)
			}
		}
	}

	for i := range messages {
		// access to each message
		msg := messages[i]

		// lock in memory repo to write message
		g.messageRepo.Mu.Lock()

		// write in memory repo key = topic , value = message and publisher
		g.messageRepo.Data[topic] = append(g.messageRepo.Data[topic], publisherMessage{m: msg, p: publisher})

		// unlock after write
		g.messageRepo.Mu.Unlock()

		// send all message
		err := g.sendMessage(publisher.name, topic, msg)
		if err != nil {
			logger.L().Info("error while sending message", "error", err)

			return err
		}
	}

	return nil
}

func (g *PubSub) sendMessage(publisherName, topic string, message *Message) error {
	// get all subscriber in topic
	subscribers := g.topicSubscribers(topic)

	// check subscriber existens if not exist any subscriber return error
	if len(subscribers) == 0 {
		logger.L().Info("no subscribers have")

		return fmt.Errorf("no subscribers have")
	}

	// i do not publisher wait to sending message so send message in other goroutin
	go func(subscribers []*subscriber) {
		// create wait group to ensur all message send in all goroutin
		wg := &sync.WaitGroup{}

		// send message to each subscribers
		for i := range subscribers {
			subscriber := subscribers[i]

			wg.Add(1)
			go func() {
				subscriber.sendMessageToSubscriber(publisherName, message)
				wg.Done()
			}()
		}

		wg.Wait()
	}(subscribers)

	return nil
}

func (g *PubSub) topicSubscribers(topic string) []*subscriber {
	subscribers, ok := g.subscribers[topic]
	if !ok {
		return nil
	}

	// let's do a copy to avoid race conditions and deadlocks due to lock
	subscribersCopy := make([]*subscriber, len(subscribers))
	copy(subscribersCopy, subscribers)

	return subscribersCopy
}

func (g *PubSub) Subscribe(topic string) (chan *Message, error) {
	g.closedLock.Lock()

	if g.closed {
		g.closedLock.Unlock()
		logger.L().Info("Pub/Sub closed")

		return nil, fmt.Errorf("Pub/Sub closed")
	}

	g.subscribersWg.Add(1)
	g.closedLock.Unlock()

	g.subscribersLock.Lock()

	topicLock, _ := g.subscribersByTopicLock.LoadOrStore(topic, &sync.Mutex{})
	tl, ok := topicLock.(*sync.Mutex)
	if !ok {
		logger.L().Info("type assertion failed in pubsub.publish")

		panic("type assertion failed in pubsub.publish")
	}
	tl.Lock()

	s := &subscriber{
		id:             uuid.New(),
		messageChannel: make(chan *Message, g.cfg.SubscriberChannelBufferSize),
	}

	defer g.subscribersLock.Unlock()
	defer tl.Unlock()

	g.addSubscriber(topic, s)

	return s.messageChannel, nil
}

func (g *PubSub) addSubscriber(topic string, s *subscriber) {
	if _, ok := g.subscribers[topic]; !ok {
		g.subscribers[topic] = make([]*subscriber, 0)
	}

	g.subscribers[topic] = append(g.subscribers[topic], s)
}

func (g *PubSub) isClosed() bool {
	g.closedLock.Lock()
	defer g.closedLock.Unlock()

	return g.closed
}

// Close closes the GoChannel Pub/Sub.
func (g *PubSub) Close() error {
	g.closedLock.Lock()
	defer g.closedLock.Unlock()

	if g.closed {
		return nil
	}

	g.closed = true

	logger.L().Info("wait to send message to all subscriber then close pubsub")

	g.subscribersWg.Wait()

	return nil
}

func (g *PubSub) RemoveReceivedMessageFromMemoryRepo(topic string) {
	// lock message repo
	g.messageRepo.Mu.Lock()
	// check existens message in topic
	_, ok := g.messageRepo.Data[topic]
	if ok {
		// if there are/is remove from repo
		delete(g.messageRepo.Data, topic)
		logger.L().Info("message received to one subscriber and delete message from memory repo")

		// unlouk message repo
		g.messageRepo.Mu.Unlock()
	}
}
