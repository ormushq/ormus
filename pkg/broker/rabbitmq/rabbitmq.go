package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	MessageBroker "github.com/ormushq/ormus/pkg/broker/messagebroker"
	"github.com/streadway/amqp"
)

// RabbitMQ represents a RabbitMQ message broker client.
type RabbitMQ struct {
	conn         *amqp.Connection
	ch           *amqp.Channel
	exchangeName string
	exchangeMode string
	wg           sync.WaitGroup // WaitGroup for synchronization
}

// NewRabbitMQBroker creates a new instance of RabbitMQ.
func NewRabbitMQBroker(amqpCfg *AMQPConfig) (*RabbitMQ, error) {
	// generate the AMQP URI from the AMQPConfig.
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		amqpCfg.Username,
		amqpCfg.Password,
		amqpCfg.Hostname,
		amqpCfg.Port,
		amqpCfg.VirtualHost,
	)
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQ{
		conn:         conn,
		ch:           ch,
		exchangeName: amqpCfg.ExchangeName,
		exchangeMode: amqpCfg.ExchangeMode,
	}, nil
}

// PublishMessage publishes messages to a specified topic in RabbitMQ.
func (rb *RabbitMQ) PublishMessage(topic string, messages ...MessageBroker.Message) error {
	queue, err := rb.DeclareExchangeAndBindQueue(topic, rb.exchangeName, rb.exchangeMode, true)
	if err != nil {
		return err
	}
	for _, msg := range messages {
		err = rb.ch.Publish(
			rb.exchangeName, // exchange
			queue.Name,      // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg.Payload,
			})
		if err != nil {
			return fmt.Errorf("failed to publish a message: %w", err)
		}
	}

	return nil
}

func (rb *RabbitMQ) CloseChannel() error {
	return rb.ch.Close()
}

// DeclareExchange declares a new exchange with the given name and type.
func (rb *RabbitMQ) DeclareExchange(exchangeName, kind string) error {
	return rb.ch.ExchangeDeclare(
		exchangeName, // name
		kind,         // type: "direct", "fanout", "topic", "headers"
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func (rb *RabbitMQ) DeclareExchangeAndBindQueue(topic, exchangeName, kind string, autoDelete bool) (*amqp.Queue, error) {
	err := rb.DeclareExchange(exchangeName, kind)
	if err != nil {
		return nil, err
	}

	return rb.DeclareAndBindQueue(topic, exchangeName, autoDelete)
}

func (rb *RabbitMQ) DeclareAndBindQueue(topic, exchangeName string, autoDelete bool) (*amqp.Queue, error) {
	q, err := rb.ch.QueueDeclare(
		topic,      // name
		false,      // durable
		autoDelete, // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to Declare a queue: %w", err)
	}
	err = rb.ch.QueueBind(
		q.Name,       // queue name
		topic,        // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind a queue: %w", err)
	}

	return &q, nil
}

// ConsumeMessage consumes messages from a specified topic in RabbitMQ.
func (rb *RabbitMQ) ConsumeMessage(topic string) (<-chan *MessageBroker.Message, error) {
	msgs, err := rb.ch.Consume(
		topic, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}
	out := make(chan *MessageBroker.Message)
	go func() {
		defer close(out) // Close the 'out' channel when all messages have been processed
		for amqpMsg := range msgs {
			rb.wg.Add(1)
			go func(amqpMsg amqp.Delivery) {
				defer rb.wg.Done()
				out <- &MessageBroker.Message{
					ID:      uuid.New(),
					Topic:   topic,
					Payload: amqpMsg.Body,
				}
			}(amqpMsg)
		}
		rb.wg.Wait() // Wait for all goroutines to finish
	}()

	return out, nil
}

// Close closes the RabbitMQ connection.
func (rb *RabbitMQ) Close() error {
	err := rb.ch.Close()
	if err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}

	err = rb.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}
