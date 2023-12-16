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

func TestPubSub_Close(t *testing.T) {

	// Given
	topic := "happy-topic"
	pubsub := observer.NewPubsub(observer.Config{SubscriberChannelBufferSize: 10})
	defer pubsub.Close()

	message := observer.NewMessage(uuid.New(), "hello, world!")

	// When

	go func() {
		_, _ = pubsub.Subscribe(topic)
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		err := pubsub.Publish(topic, message)
		require.NoError(t, err)
	}()

	pubsub.Close()

	// Then
	assert.Panics(t, func() {
		pubsub.Publish(topic, message)
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

	err := pubsub.Publish("test", message)
	if err != nil {
		t.Errorf("Failed to publish message: %v", err)
		return
	}

	// Wait for all goroutines to finish
	wg.Wait()

}
