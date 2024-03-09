package rabbitmq

import (
	"fmt"
	"github.com/ormushq/ormus/logger"
	"testing"
	"time"

	"github.com/ormushq/ormus/pkg/broker/messagebroker"
	"github.com/ormushq/ormus/pkg/broker/rabbitmq"
)

// Define a struct to hold parameters for the fanout test case
type FanOutTestCase struct {
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
	testCases := []FanOutTestCase{
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

func runFanoutTest(t *testing.T, tc FanOutTestCase) {
	conn := make(map[int]*rabbitmq.RabbitMQ)
	for i := 0; i < len(tc.QueueNames); i++ {
		cfg := rabbitmq.AMQPConfig{
			Username:     "guest",
			Password:     "guest",
			Hostname:     "localhost",
			Port:         5672,
			VirtualHost:  "/",
			ExchangeName: tc.ExchangeName[i],
			ExchangeMode: tc.Mode,
		}
		conn[i] = setupRabbitMQFanout(t, &cfg)

	}
	defer deferAllConnfan(conn, t)

	// Publish messages to the fanout exchange
	for i := range conn {
		publishMessagesFanout(t, conn[i], tc.QueueNames[i], tc.NumMessages)
	}

	// Consume messages from each queue and verify counts
	checkMessagesReceivedFanout(t, conn, tc)
}

func deferAllConnfan(conn map[int]*rabbitmq.RabbitMQ, t *testing.T) {
	for i := 0; i < len(conn); i++ {
		if err := conn[i].Close(); err != nil {
			t.Fatalf("Failed to close RabbitMQ connection: %v", err)
		}
	}
}

// Helper function to create a RabbitMQ connection
func setupRabbitMQFanout(t *testing.T, cfg *rabbitmq.AMQPConfig) *rabbitmq.RabbitMQ {
	amqpCfg := rabbitmq.NEWAMQPConfig(cfg)
	conn, err := rabbitmq.NewRabbitMQBroker(amqpCfg)
	if err != nil {
		t.Fatalf("Failed to create RabbitMQ broker: %v", err)
	}
	return conn
}

func publishMessagesFanout(t *testing.T, conn *rabbitmq.RabbitMQ, topic string, numMessages int) {
	for i := 0; i < numMessages; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		err := conn.PublishMessage(topic, messagebroker.NewMessage(topic, []byte(message)))
		if err != nil {
			t.Fatalf("Failed to publish message to exchange %s: %v", topic, err)
		}
	}
}

// Helper function to consume messages from queues and verify counts
func checkMessagesReceivedFanout(t *testing.T, conns map[int]*rabbitmq.RabbitMQ, tc FanOutTestCase) {
	receivedCount := 0

	// Create a map to track the number of messages received from each queue
	receivedMap := make(map[string]int)

	// Create channels for each queue
	channels := make([]<-chan *messagebroker.Message, len(tc.QueueNames))
	for i, conn := range conns {
		chMsg, err := conn.ConsumeMessage(tc.QueueNames[i])
		if err != nil {
			t.Fatalf("Failed to start consuming messages from queue %s: %v", tc.QueueNames[i], err)
		}
		channels[i] = chMsg
		receivedMap[tc.QueueNames[i]] = 0
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
				default:
					// Do nothing, move to the next channel
				}
			}
			// Check if all expected messages are received
			if receivedCount > tc.Expected {
				t.Fatalf("err")
			}
			// Check if all expected messages are received
			if receivedCount == tc.Expected {
				logger.L().Debug("Received all expected messages.")
				return
			}
		}
	}
}
