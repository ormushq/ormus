package test

import (
	"fmt"
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
			NumMessages:  10000,
			ExpectedCounts: map[string]int{
				"queue1": 10000,
				"queue2": 10000,
				"queue3": 10000,
				"queue4": 10000,
			},
			//! expected = number of message * number of queue = 400 => for all queue :if auto-delete be false
			//! and the queue exist before =1600 and 1000 if auto-delete be true
			Expected: 100000,
		},
		{
			Name:         "different exchange",
			Mode:         "fanout",
			ExchangeName: []string{"fanout_exchange1", "fanout_exchange2", "fanout_exchange3", "fanout_exchange4"},
			QueueNames:   []string{"queue1", "queue2", "queue3", "queue4"},
			NumMessages:  10000,
			ExpectedCounts: map[string]int{
				"queue1": 10000,
				"queue2": 10000,
				"queue3": 10000,
				"queue4": 10000,
			},
			Expected: 40000,
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
		cfg := rabbitmq.AMQPBaseConfig{
			Username:     "guest",
			Password:     "guest",
			Hostname:     "localhost",
			Port:         5672,
			VirtualHost:  "/",
			ExchangeName: tc.ExchangeName[i],
			ExchangeMode: tc.Mode,
		}
		conn[i] = setupRabbitMQFanout(t, cfg)

	}
	defer deferAllConnfan(conn, t)

	// Publish messages to the fanout exchange
	for i := range conn {
		publishMessagesFanout(t, conn[i], tc.QueueNames[i], tc.NumMessages)
	}
	// sleep 5 second to wait for see the published messages on ui
	// time.Sleep(5 * time.Second)
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
func setupRabbitMQFanout(t *testing.T, cfg rabbitmq.AMQPBaseConfig) *rabbitmq.RabbitMQ {
	amqpCfg := rabbitmq.NEWAMQPConfig(cfg, nil)
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
			allClosed := true
			for i, ch := range channels {
				select {
				case _, ok := <-ch:
					if !ok {
						continue
					}
					receivedCount++
					receivedMap[tc.QueueNames[i]]++
					allClosed = false
				default:
					// Do nothing, move to the next channel
				}
			}
			if allClosed {
				break
			}
		}
		if receivedCount == tc.Expected {
			break
		}

	}
	fmt.Printf("Received %d messages, expected %d\n", receivedCount, tc.Expected)
}
