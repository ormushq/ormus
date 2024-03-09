package rabbitmq

import (
	"fmt"
	"testing"
	"time"

	"github.com/ormushq/ormus/pkg/broker/messagebroker"
	"github.com/ormushq/ormus/pkg/broker/rabbitmq"
)

// Define a TestCase struct to hold parameters for each test case
type DirectTestCase struct {
	Name         string
	ExchangeName string
	Kind         string
	NumWorkers   int
	NumMessages  int
	ExpectedMsg  int
}

func TestRabbitMQConcurrentConsumption(t *testing.T) {
	// Define test cases
	directTestCases := []DirectTestCase{
		{
			Name:         "SingleWorker",
			ExchangeName: "test_exchange",
			Kind:         "direct",
			NumWorkers:   1,
			NumMessages:  10,
			ExpectedMsg:  10,
		},
		{
			Name:         "MultipleWorkers",
			ExchangeName: "test_exchange",
			Kind:         "direct",
			NumWorkers:   10,
			NumMessages:  100,
			ExpectedMsg:  1000,
		},
	}
	// Run test cases
	for _, tc := range directTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			runTest(t, tc)
			time.Sleep(10 * time.Second)
		})
	}
}

func runTest(t *testing.T, tc DirectTestCase) {
	conn := make(map[int]*rabbitmq.RabbitMQ)
	queueName := "test_queue"
	for i := 0; i < tc.NumWorkers; i++ {
		conn[i] = setupRabbitMQDir(t, tc)
		// Publish messages
		publishMessagesDir(t, conn[i], queueName, tc.NumMessages)
	}

	defer deferAllConn(conn, t)
	// Start worker goroutines
	channels := startWorkers(t, conn, queueName, tc.NumWorkers)
	// Wait for messages to be consumed
	time.Sleep(500 * time.Millisecond)
	// Check if the correct number of messages was received
	checkMessagesReceivedDir(t, channels, tc.ExpectedMsg)
}

func deferAllConn(conn map[int]*rabbitmq.RabbitMQ, t *testing.T) {
	for i := 0; i < len(conn); i++ {
		if err := conn[i].Close(); err != nil {
			t.Fatalf("Failed to close RabbitMQ connection: %v", err)
		}
	}
}

func setupRabbitMQDir(t *testing.T, tc DirectTestCase) *rabbitmq.RabbitMQ {
	cfg := rabbitmq.AMQPConfig{
		Username:     "guest",
		Password:     "guest",
		Hostname:     "localhost",
		Port:         5672,
		VirtualHost:  "/",
		ExchangeName: tc.ExchangeName,
		ExchangeMode: tc.Kind,
	}
	amqpCfg := rabbitmq.NEWAMQPConfig(&cfg)
	conn, err := rabbitmq.NewRabbitMQBroker(amqpCfg)
	if err != nil {
		t.Fatalf("Failed to create RabbitMQ broker: %v", err)
	}
	return conn
}

func startWorkers(t *testing.T, conn map[int]*rabbitmq.RabbitMQ, queueName string, numWorkers int) []<-chan *messagebroker.Message {
	channels := make([]<-chan *messagebroker.Message, numWorkers)
	for i := 0; i < len(conn); i++ {
		chMsg, err := conn[i].ConsumeMessage(queueName)
		if err != nil {
			t.Fatalf("Failed to start worker: %v", err)
		}
		channels[i] = chMsg
	}

	return channels
}

func publishMessagesDir(t *testing.T, conn *rabbitmq.RabbitMQ, topic string, numMessages int) {
	for i := 0; i < numMessages; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		err := conn.PublishMessage(topic, messagebroker.NewMessage(topic, []byte(message)))
		if err != nil {
			t.Fatalf("Failed to publish message: %v", err)
		}
	}
}

func checkMessagesReceivedDir(t *testing.T, channels []<-chan *messagebroker.Message, expected int) {
	received := 0
	for received < expected {
		select {
		case <-time.After(1 * time.Second):
			t.Fatalf("Timeout: Received %d messages, expected %d", received, expected)
		default:
			for _, ch := range channels {
				select {
				case _, ok := <-ch:
					if !ok {
						continue
					}
					received++

				default:
					// Do nothing, move to the next channel
				}
			}
		}
	}
	if received != expected {
		t.Errorf("Received %d messages, expected %d", received, expected)
	}
}
