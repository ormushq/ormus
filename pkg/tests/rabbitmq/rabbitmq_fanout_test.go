package rabbitmq

import (
	"fmt"
	"github.com/ormushq/ormus/pkg/broker/message_broker"
	"github.com/ormushq/ormus/pkg/broker/rabbitmq"
	"testing"
	"time"
)

// Define a struct to hold parameters for the fanout test case
type FanoutTestCase struct {
	Name           string
	Mode           string
	ExchangeName   []string
	QueueNames     []string
	NumMessages    int
	ExpectedCounts map[string]int
	Expected       int
}

func TestFanoutMessaging(t *testing.T) {
	// Define test cases
	testCases := []FanoutTestCase{
		{
			Name:         "same exchange",
			Mode:         "fanout",
			ExchangeName: []string{"fanout_exchange", "fanout_exchange", "fanout_exchange", "fanout_exchange"},
			QueueNames:   []string{"queue1", "queue2", "queue3", "queue4"},
			NumMessages:  100,
			ExpectedCounts: map[string]int{
				"queue1": 100,
				"queue2": 100,
				"queue3": 100,
				"queue4": 100,
			},
			//! expected = number of message * number of queue = 400 => for all queue :if auto-delete be false
			//! and the queue exist before =1600 and 1000 if auto-delete be true
			Expected: 1000,
		},
		{
			Name:         "different exchange",
			Mode:         "fanout",
			ExchangeName: []string{"fanout_exchange1", "fanout_exchange2", "fanout_exchange3", "fanout_exchange4"},
			QueueNames:   []string{"queue1", "queue2", "queue3", "queue4"},
			NumMessages:  100,
			ExpectedCounts: map[string]int{
				"queue1": 100,
				"queue2": 100,
				"queue3": 100,
				"queue4": 100,
			},
			Expected: 400,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			runFanoutTest(t, tc)
			time.Sleep(10 * time.Second)
		})
	}
}

func runFanoutTest(t *testing.T, tc FanoutTestCase) {
	// Create a RabbitMQ connection
	conn := setupRabbitMQFanout(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatalf("Failed to close RabbitMQ connection: %v", err)
		}
	}()

	// Publish messages to the fanout exchange
	for i := range tc.QueueNames {
		// Declare the fanout exchange
		err := conn.DeclareExchange(tc.ExchangeName[i], tc.Mode)
		if err != nil {
			t.Fatalf("Failed to declare fanout exchange %s: %v", tc.ExchangeName, err)
		}
		err = DeclareAndBindQueueFanout(t, conn, tc.QueueNames[i], tc.ExchangeName[i], true)
		if err != nil {
			t.Fatalf("Failed to create Queue: %v", err)
		}
		publishMessagesFanout(t, conn, tc.QueueNames[i], tc.NumMessages)
	}

	// Consume messages from each queue and verify counts
	checkMessagesReceivedFanout(t, conn, tc)
}
func DeclareAndBindQueueFanout(t *testing.T, conn *rabbitmq.RabbitMQ, topic, ExchangeName string, autoDelete bool) error {
	q, err := conn.DeclareAndBindQueue(topic, ExchangeName, autoDelete)
	if err != nil {
		return err
	}
	fmt.Println("Queue created :", q.Name)
	return nil
}

// Helper function to create a RabbitMQ connection
func setupRabbitMQFanout(t *testing.T) *rabbitmq.RabbitMQ {
	amqpCfg := rabbitmq.DefaultAMQPConfig()
	conn, err := rabbitmq.NewRabbitMQBroker(amqpCfg)
	if err != nil {
		t.Fatalf("Failed to create RabbitMQ broker: %v", err)
	}
	return conn
}

// Helper function to publish messages to an exchange
func publishMessagesFanout(t *testing.T, conn *rabbitmq.RabbitMQ, topic string, numMessages int) {
	for i := 0; i < numMessages; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		time.Sleep(125 * time.Millisecond)
		err := conn.PublishMessage(topic, message_broker.NewMessage(topic, []byte(message)))
		if err != nil {
			t.Fatalf("Failed to publish message to exchange %s: %v", topic, err)
		}
	}
}

// Helper function to consume messages from a queue and verify counts
// Helper function to consume messages from queues and verify counts
func checkMessagesReceivedFanout(t *testing.T, conn *rabbitmq.RabbitMQ, tc FanoutTestCase) {
	receivedCount := 0

	// Create a map to track the number of messages received from each queue
	receivedMap := make(map[string]int)

	// Create channels for each queue
	channels := make([]<-chan *message_broker.Message, len(tc.QueueNames))
	for i, queueName := range tc.QueueNames {
		time.Sleep(125 * time.Millisecond)
		conn.SetExchangeName(tc.ExchangeName[i])
		chmsg, err := conn.ConsumeMessage(queueName)
		if err != nil {
			t.Fatalf("Failed to start consuming messages from queue %s: %v", queueName, err)
		}
		channels[i] = chmsg
		receivedMap[queueName] = 0
	}

	// Use a select statement to receive messages from all queues
	for {
		select {
		// Receive messages from each queue
		case <-time.After(1 * time.Second):

			t.Fatalf("Timeout: Received %d messages, expected %d", receivedCount, tc.Expected)
		default:
			for i, ch := range channels {
				select {
				case _, ok := <-ch:
					if !ok {
						continue
					}
					receivedCount++
					receivedMap[tc.QueueNames[i]]++
					fmt.Println("received:", receivedCount)
				default:
					// Do nothing, move to the next channel
				}
			}
			// Check if all expected messages are received
			if receivedCount > tc.Expected {
				t.Fatalf("err")
			}
			fmt.Println("received:", receivedCount, "Expected:", tc.Expected)
			// Check if all expected messages are received
			if receivedCount == tc.Expected {
				fmt.Println("Received all expected messages.")
				return
			}
		}
	}
}