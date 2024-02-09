package rabbitmq

import (
	"fmt"
	"github.com/ormushq/ormus/pkg/broker/message_broker"
	"github.com/ormushq/ormus/pkg/broker/rabbitmq"
	"testing"
	"time"
)

// Define a TestCase struct to hold parameters for each test case
type DirectTestCase struct {
	Name         string
	NumWorkers   int
	NumMessages  int
	ExpectedMsgs int
}

func TestRabbitMQConcurrentConsumption(t *testing.T) {
	// Define test cases
	directtestCases := []DirectTestCase{
		{
			Name:         "SingleWorker",
			NumWorkers:   1,
			NumMessages:  10,
			ExpectedMsgs: 10,
		},
		{
			Name:         "MultipleWorkers",
			NumWorkers:   10,
			NumMessages:  100,
			ExpectedMsgs: 100,
		},
	}

	// Run test cases
	for _, tc := range directtestCases {
		t.Run(tc.Name, func(t *testing.T) {
			runTest(t, tc)
		})
	}
}

func runTest(t *testing.T, tc DirectTestCase) {
	conn := setupRabbitMQDir(t)
	defer func() {
		if err := conn.Close(); err != nil {
			t.Fatalf("Failed to close RabbitMQ connection: %v", err)
		}
	}()
	queueName := "test_queue"

	// Publish messages
	publishMessagesDir(t, conn, queueName, tc.NumMessages)

	// Start worker goroutines
	channels := startWorkers(t, conn, queueName, tc.NumWorkers)

	// Wait for messages to be consumed
	time.Sleep(500 * time.Millisecond)

	// Check if the correct number of messages was received
	checkMessagesReceivedDir(t, channels, tc.ExpectedMsgs)
}

func setupRabbitMQDir(t *testing.T) *rabbitmq.RabbitMQ {
	amqpCfg := rabbitmq.DefaultAMQPConfig()
	conn, err := rabbitmq.NewRabbitMQBroker(amqpCfg)
	if err != nil {
		t.Fatalf("Failed to create RabbitMQ broker: %v", err)
	}

	// Declare exchange
	err = conn.DeclareExchange("test_exchange", "direct")
	if err != nil {
		t.Fatalf("Failed to declare exchange: %v", err)
	}

	return conn
}

func startWorkers(t *testing.T, conn *rabbitmq.RabbitMQ, queueName string, numWorkers int) []<-chan *message_broker.Message {
	channels := make([]<-chan *message_broker.Message, numWorkers)

	for i := 0; i < numWorkers; i++ {
		chmsg, err := conn.ConsumeMessage(queueName)
		if err != nil {
			t.Fatalf("Failed to start worker: %v", err)
		}
		channels[i] = chmsg
	}

	return channels
}

func publishMessagesDir(t *testing.T, conn *rabbitmq.RabbitMQ, topic string, numMessages int) {
	for i := 0; i < numMessages; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		err := conn.PublishMessage(topic, message_broker.NewMessage(topic, []byte(message)))
		if err != nil {
			t.Fatalf("Failed to publish message: %v", err)
		}
	}

}
func checkMessagesReceivedDir(t *testing.T, channels []<-chan *message_broker.Message, expected int) {

	received := 0
	for received < expected {
		select {
		case <-time.After(1 * time.Second):
			t.Fatalf("Timeout: Received %d messages, expected %d", received, expected)
		default:
			for _, ch := range channels {
				select {
				case msg, ok := <-ch:
					if !ok {
						continue
					}
					received++
					fmt.Printf("ID : %v , topic : %v , payload : %s received : %d \n", msg.ID, msg.Topic, string(msg.Payload), received)
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
