package pubsub_test

import (
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/pkg/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPublishSubscribe(t *testing.T) {
	var wg sync.WaitGroup

	newPubSub := pubsub.NewPubSub(pubsub.Config{
		SubscriberChannelBufferSize: 100,
	},
	)

	message := pubsub.NewMessage(uuid.New(), "hello", *pubsub.NewPublisher("publisher"))

	wg.Add(1)
	go func(PubSub *pubsub.PubSub) {
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
	go func(PubSub *pubsub.PubSub) {
		time.Sleep(100 * time.Millisecond)
		p := pubsub.NewPublisher("publisher")
		err := PubSub.Publish(p, "test", message)
		assert.NoError(t, err)
		slog.Info("message published")
		wg.Done()
	}(newPubSub)

	wg.Wait()
}

func TestPubSub_Close(t *testing.T) {
	// Given
	topic := "happy-topic"
	Gopubsub := pubsub.NewPubSub(pubsub.Config{SubscriberChannelBufferSize: 10})
	defer Gopubsub.Close()

	message := pubsub.NewMessage(uuid.New(), "hello, world!", *pubsub.NewPublisher("publisher"))

	// When

	go func() {
		_, _ = Gopubsub.Subscribe(topic)
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		p := pubsub.NewPublisher("publisher")
		err := Gopubsub.Publish(p, topic, message)
		require.NoError(t, err)
	}()

	Gopubsub.Close()

	// Then
	assert.Panics(t, func() {
		p := pubsub.NewPublisher("publisher")
		Gopubsub.Publish(p, topic, message)
	})
}

func TestPubsub_Subscribe_RaceCondition(t *testing.T) {
	config := pubsub.Config{SubscriberChannelBufferSize: 10}
	Gopubsub := pubsub.NewPubSub(config)

	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup

	// Number of goroutines to create
	numGoroutines := 100

	message := pubsub.NewMessage(uuid.New(), "Race condition test message", *pubsub.NewPublisher("publisher"))

	// Create multiple goroutines trying to subscribe simultaneously
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Random topic for each goroutine
			topic := "test"

			subscriber, err := Gopubsub.Subscribe(topic)
			if err != nil {
				t.Errorf("Failed to create subscriber: %v", err)
				return
			}

			// Receive and verify the message
			select {
			case receivedMessage := <-subscriber:
				if receivedMessage.ID != message.ID || receivedMessage.Payload != message.Payload {
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
	p := pubsub.NewPublisher("publisher")
	err := Gopubsub.Publish(p, "test", message)
	if err != nil {
		t.Errorf("Failed to publish message: %v", err)
		return
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
