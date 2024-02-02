package pubsub

import (
	"errors"
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

	messageRepo     map[string][]InMemoryMessageStore
	messageRepoLock sync.Mutex
}

func NewPubSub(cfg Config) *PubSub {
	return &PubSub{
		cfg:                    cfg,
		subscribers:            make(map[string][]*subscriber),
		subscribersByTopicLock: sync.Map{},
		messageRepo:            make(map[string][]InMemoryMessageStore, 0),
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

	g.messageRepoLock.Lock()
	// add unsended message to messages
	unsendedMessages, ok := g.messageRepo[topic]
	// if there are messages ok == true
	if ok {
		for _, msg := range unsendedMessages {
			// then we need check did this publisher send this message or not?
			if msg.pubslisher == publisher {
				// append to new messages list to send again
				messages = append(messages, msg.message)
			}
		}
	}

	g.messageRepoLock.Unlock()

	subscriber := g.topicSubscribers(topic)

	// add new message to repo
	for _, msg := range messages {

		g.messageRepoLock.Lock()
		unsendedMessages := g.messageRepo[topic]

		for _, unmsg := range unsendedMessages {
			// check if not duplicate add to repo
			if unmsg.message != msg {
				unsendedMessages = append(unsendedMessages, InMemoryMessageStore{message: msg, pubslisher: publisher, subscriberCount: len(subscriber)})
			}
		}

		g.messageRepoLock.Unlock()
	}

	for i := range messages {
		// access to each message
		msg := messages[i]

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

func (g *PubSub) RemoveReceivedMessageFromMemoryRepo(publisher *Publisher, topic string, message *Message) error {
	// lock message repo
	g.messageRepoLock.Lock()
	// check existens message in topic
	pms, ok := g.messageRepo[topic]
	if !ok {
		logger.L().Info("there is no message on this topic")

		// unlouk message repo
		g.messageRepoLock.Unlock()

		return errors.New("there is no message on this topic")
	}

	for _, pm := range pms {

		if pm.pubslisher == publisher {
			if pm.message == message {
				pm.subscriberCount--
			}
		}

		if pm.subscriberCount == 0 {
			pms = removeFromMemoryStorge(pms, publisher, message)
		}
	}

	g.messageRepo[topic] = pms

	g.messageRepoLock.Unlock()

	return nil
}

func removeFromMemoryStorge(datas []InMemoryMessageStore, publisher *Publisher, message *Message) []InMemoryMessageStore {
	indexToRemove := -1
	for i, pm := range datas {
		if pm.pubslisher == publisher {
			if pm.message == message {
				indexToRemove = i
			}
		}
	}

	// If the element is found, remove it
	if indexToRemove != -1 {
		datas = append(datas[:indexToRemove], datas[indexToRemove+1:]...)
	}

	return datas
}
