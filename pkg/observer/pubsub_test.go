package observer_test

import (
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/pkg/observer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPublishSubscribe(t *testing.T) {

	var wg sync.WaitGroup

	newPubSub := observer.NewPubsub(observer.Config{
		SubscriberChannelBufferSize: 100},
	)

	message := observer.NewMessage(uuid.New(), "hello")

	wg.Add(1)
	go func(PubSub *observer.GoChannelPubsub) {

		msg, err := PubSub.Subscribe("test")
		assert.NoError(t, err)

		go func() {
		Loop:
			select {
			case receivedMessage := <-msg:
				slog.Info("message", "msg", receivedMessage)
				assert.Equal(t, message, receivedMessage)
				wg.Done()
				return
			default:
				goto Loop
			}
		}()

	}(newPubSub)

	wg.Add(1)
	go func(PubSub *observer.GoChannelPubsub) {

		time.Sleep(100 * time.Millisecond)
		err := PubSub.Publish("test", message)
		assert.NoError(t, err)
		slog.Info("message published")
		wg.Done()

	}(newPubSub)

	wg.Wait()
}

func TestPubsub_Publish_Subscribe(t *testing.T) {

	t.Run("subscriber receives multiple messages", func(t *testing.T) {

		// Given
		topic := "happy-topic"
		pubsub := observer.NewPubsub(observer.Config{SubscriberChannelBufferSize: 10})
		defer pubsub.Close()

		messages := []*observer.Message{
			observer.NewMessage(uuid.New(), "hello, world 1!"),
			observer.NewMessage(uuid.New(), "hello, world 2!"),
			observer.NewMessage(uuid.New(), "hello, world 3!"),
		}

		// When
		err := pubsub.Publish(topic, messages...)
		require.NoError(t, err)

		subscriberChannel, err := pubsub.Subscribe(topic)
		require.NoError(t, err)

		// Then
		var receivedMessages []*observer.Message
		for range messages {
			receivedMessages = append(receivedMessages, <-subscriberChannel)
		}

		assert.ElementsMatch(t, messages, receivedMessages)
	})

	t.Run("multiple subscribers receive the message", func(t *testing.T) {

		// Given
		topic := "happy-topic"
		pubsub := observer.NewPubsub(observer.Config{SubscriberChannelBufferSize: 10})
		defer pubsub.Close()

		message := observer.NewMessage(uuid.New(), "hello, world!")

		// When
		err := pubsub.Publish(topic, message)
		require.NoError(t, err)

		subscriberChannels := make([]chan *observer.Message, 5)
		for i := range subscriberChannels {
			subscriberChannels[i], err = pubsub.Subscribe(topic)
			require.NoError(t, err)
		}

		// Then
		var receivedMessages []*observer.Message
		for i := range subscriberChannels {
			receivedMessages = append(receivedMessages, <-subscriberChannels[i])
		}

		for _, receivedMessage := range receivedMessages {
			assert.Equal(t, message.Id, receivedMessage.Id)
			assert.Equal(t, message.Payload, receivedMessage.Payload)
		}
	})
}

func TestPubSub_Close(t *testing.T) {

	t.Run("pubsub closes properly", func(t *testing.T) {

		// Given
		topic := "happy-topic"
		pubsub := observer.NewPubsub(observer.Config{SubscriberChannelBufferSize: 10})
		defer pubsub.Close()

		message := observer.NewMessage(uuid.New(), "hello, world!")

		// When
		err := pubsub.Publish(topic, message)
		require.NoError(t, err)

		pubsub.Close()

		// Then
		assert.Panics(t, func() {
			pubsub.Publish(topic, message)
		})
	})
}

func TestPubsub_Subscribe_RaceCondition(t *testing.T) {

	config := observer.Config{SubscriberChannelBufferSize: 10}
	pubsub := observer.NewPubsub(config)

	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup

	// Number of goroutines to create
	numGoroutines := 100

	message := observer.NewMessage(uuid.New(), "Race condition test message")

	// Create multiple goroutines trying to subscribe simultaneously
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Random topic for each goroutine
			topic := "test"

			subscriber, err := pubsub.Subscribe(topic)
			if err != nil {
				t.Errorf("Failed to create subscriber: %v", err)
				return
			}

			// Receive and verify the message
			select {
			case receivedMessage := <-subscriber:
				if receivedMessage.Id != message.Id || receivedMessage.Payload != message.Payload {
					t.Errorf("Received unexpected message. Expected: %+v, Received: %+v", message, receivedMessage)
				}
			case <-time.After(time.Millisecond * 500):
				t.Error("Timed out waiting for message.")
			}
		}()

	}

	// Simulate some processing time
	time.Sleep(time.Millisecond * 100)

	// Publish a message to the topic

	err := pubsub.Publish("test", message)
	if err != nil {
		t.Errorf("Failed to publish message: %v", err)
		return
	}

	// Wait for all goroutines to finish
	wg.Wait()

}
