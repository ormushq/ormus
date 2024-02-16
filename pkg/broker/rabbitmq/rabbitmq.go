package rabbitmq

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	MessageBroker "github.com/ormushq/ormus/pkg/broker/messagebroker"
	"github.com/streadway/amqp"
)

const sleepTime = 125

// RabbitMQ represents a RabbitMQ message broker client.
type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
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
		conn: conn,
		ch:   ch,
	}, nil
}

// PublishMessage publishes messages to a specified topic in RabbitMQ.
func (rb *RabbitMQ) PublishMessage(topic string, exchangeName string, messages ...*MessageBroker.Message) error {
	time.Sleep(sleepTime * time.Millisecond)
	for _, msg := range messages {
		err := rb.ch.Publish(
			exchangeName, // exchange
			topic,        // routing key
			false,        // mandatory
			false,        // immediate
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

// ConsumeMessage consumes messages from a specified topic in RabbitMQ.
func (rb *RabbitMQ) ConsumeMessage(topic string) (<-chan *MessageBroker.Message, error) {
	time.Sleep(sleepTime * time.Millisecond)
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
		for amqpMsg := range msgs {
			out <- &MessageBroker.Message{
				ID:      uuid.New(),
				Topic:   topic,
				Payload: amqpMsg.Body,
			}
		}
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
