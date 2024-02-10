package rabbitmq

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ormushq/ormus/logger"
	MessageBroker "github.com/ormushq/ormus/pkg/broker/message_broker"
	"github.com/streadway/amqp"
	"time"
)

type additionalFields struct {
	ExchangeName string
}

// RabbitMQ represents a RabbitMQ message broker client.
type RabbitMQ struct {
	conn             *amqp.Connection
	ch               *amqp.Channel
	additionalFields additionalFields
}

func (rb *RabbitMQ) SetExchangeName(ExchangeName string) {
	rb.additionalFields.ExchangeName = ExchangeName
}
func (rb *RabbitMQ) getExchangeName() string {
	return rb.additionalFields.ExchangeName
}

func (rb *RabbitMQ) CloseChannel() error {
	return rb.ch.Close()
}

// DeclareExchange declares a new exchange with the given name and type.
func (rb *RabbitMQ) DeclareExchange(ExchangeName, kind string) error {
	// set the latest ExchangeName
	rb.additionalFields.ExchangeName = ExchangeName
	return rb.ch.ExchangeDeclare(
		ExchangeName, // name
		kind,         // type: "direct", "fanout", "topic", "headers"
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}
func (rb *RabbitMQ) DeclareAndBindQueue(topic, ExchangeName string, autoDelete bool) (*amqp.Queue, error) {
	q, err := rb.ch.QueueDeclare(
		topic,      // name
		false,      // durable
		autoDelete, // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to Declare a queue: %s", err)
	}
	logger.L().Info("queue created:", q)
	err = rb.ch.QueueBind(
		q.Name,       // queue name
		topic,        // routing key
		ExchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind a queue: %s", err)
	}
	return &q, nil

}

// NewRabbitMQBroker creates a new instance of RabbitMQ.
func NewRabbitMQBroker(amqpCfg *AMQPConfig) (*RabbitMQ, error) {
	//generate the AMQP URI from the AMQPConfig.
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		amqpCfg.Username,
		amqpCfg.Password,
		amqpCfg.Hostname,
		amqpCfg.Port,
		amqpCfg.VirtualHost,
	)
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %s", err)
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

// PublishMessage publishes messages to a specified topic in RabbitMQ.
func (rb *RabbitMQ) PublishMessage(topic string, messages ...*MessageBroker.Message) error {
	time.Sleep(125 * time.Millisecond)
	for _, msg := range messages {
		err := rb.ch.Publish(
			rb.getExchangeName(), // exchange
			topic,                // routing key
			false,                // mandatory
			false,                // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg.Payload,
			})
		if err != nil {
			return fmt.Errorf("failed to publish a message: %s", err)
		}
	}

	return nil
}

// ConsumeMessage consumes messages from a specified topic in RabbitMQ.
func (rb *RabbitMQ) ConsumeMessage(topic string) (<-chan *MessageBroker.Message, error) {
	time.Sleep(125 * time.Millisecond)
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
		return nil, fmt.Errorf("failed to register a consumer: %s", err)
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
		return fmt.Errorf("failed to close channel: %s", err)
	}

	err = rb.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %s", err)
	}

	return nil
}
