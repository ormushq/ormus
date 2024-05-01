package rabbitmq

import (
	"fmt"

	"github.com/google/uuid"
	MessageBroker "github.com/ormushq/ormus/pkg/broker/messagebroker"
	"github.com/streadway/amqp"
)

// RabbitMQ represents a RabbitMQ message broker client.
type RabbitMQ struct {
	conn        *amqp.Connection
	ch          *amqp.Channel
	BaseCfg     AMQPBaseConfig
	OptionalCfg AMQPOptions
}

// NewRabbitMQBroker creates a new instance of RabbitMQ.
func NewRabbitMQBroker(amqpCfg AMQPConfig) (*RabbitMQ, error) {
	// generate the AMQP URI from the AMQPConfig.
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		amqpCfg.baseConfig.Username,
		amqpCfg.baseConfig.Password,
		amqpCfg.baseConfig.Hostname,
		amqpCfg.baseConfig.Port,
		amqpCfg.baseConfig.VirtualHost,
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
		conn:        conn,
		ch:          ch,
		BaseCfg:     amqpCfg.baseConfig,
		OptionalCfg: amqpCfg.AMQPOption,
	}, nil
}

// PublishMessage publishes messages to a specified topic in RabbitMQ.
func (rb *RabbitMQ) PublishMessage(topic string, messages ...MessageBroker.Message) error {
	queue, err := rb.DeclareExchangeAndBindQueue(topic, rb.BaseCfg.ExchangeName, rb.BaseCfg.ExchangeMode)
	if err != nil {
		return err
	}
	// TODO : What should we do if the message is not published?
	for _, msg := range messages {
		err = rb.ch.Publish(
			rb.BaseCfg.ExchangeName,         // exchange
			queue.Name,                      // routing key
			rb.OptionalCfg.PublishMandatory, // mandatory
			rb.OptionalCfg.PublishImmediate, // immediate
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
		exchangeName,                       // name
		kind,                               // type: "direct", "fanout", "topic", "headers"
		rb.OptionalCfg.ExchangeDurable,     // durable
		rb.OptionalCfg.ExchangeAutoDeleted, // auto-deleted
		rb.OptionalCfg.ExchangeInternal,    // internal
		rb.OptionalCfg.ExchangeNoWait,      // no-wait
		rb.OptionalCfg.ExchangeArgs,        // arguments
	)
}

// DeclareExchangeAndBindQueue declares a queue to hold messages and deliver to consumers.
// Declaring creates a queue if it doesn't already exist, or ensures that an
// existing queue matches the same parameters.
func (rb *RabbitMQ) DeclareExchangeAndBindQueue(topic, exchangeName, kind string) (*amqp.Queue, error) {
	err := rb.DeclareExchange(exchangeName, kind)
	if err != nil {
		return nil, err
	}

	return rb.DeclareAndBindQueue(topic)
}

func (rb *RabbitMQ) DeclareAndBindQueue(topic string) (*amqp.Queue, error) {
	q, err := rb.ch.QueueDeclare(
		topic,                           // name
		rb.OptionalCfg.QueueDurable,     // durable
		rb.OptionalCfg.QueueAutoDeleted, // delete when unused
		rb.OptionalCfg.QueueExclusive,   // exclusive
		rb.OptionalCfg.QueueNoWait,      // no-wait
		rb.OptionalCfg.QueueArgs,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to Declare a queue: %w", err)
	}
	err = rb.ch.QueueBind(
		q.Name,                     // queue name
		topic,                      // routing key
		rb.BaseCfg.ExchangeName,    // exchange
		rb.OptionalCfg.QueueNoWait, // no-wait
		rb.OptionalCfg.QueueArgs,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind a queue: %w", err)
	}

	return &q, nil
}

// ConsumeMessage consumes messages from a specified topic in RabbitMQ.
func (rb *RabbitMQ) ConsumeMessage(topic string) (<-chan *MessageBroker.Message, error) {
	msgs, err := rb.ch.Consume(
		topic,                            // queue
		rb.OptionalCfg.ConsumerTag,       // consumer
		rb.OptionalCfg.ConsumerAutoAck,   // auto-ack
		rb.OptionalCfg.ConsumerExclusive, // exclusive
		rb.OptionalCfg.ConsumerNoLocal,   // no-local
		rb.OptionalCfg.ConsumerNoWait,    // no-wait
		rb.OptionalCfg.ConsumerArgs,      // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}
	out := make(chan *MessageBroker.Message, rb.OptionalCfg.ConsumeMessageChanSize)
	go func() {
		defer close(out) // Close the channel when the processing goroutine exits

		for amqpMsg := range msgs {
			msg := &MessageBroker.Message{
				ID:      uuid.New(),
				Topic:   topic,
				Payload: amqpMsg.Body,
			}
			out <- msg
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
