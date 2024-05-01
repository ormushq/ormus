package rabbitmq

// AMQPConfig holds configuration options for connecting to RabbitMQ using AMQP URI.
type AMQPConfig struct {
	baseConfig AMQPBaseConfig
	AMQPOption AMQPOptions
}
type AMQPBaseConfig struct {
	Username     string
	Password     string
	Hostname     string
	Port         int
	VirtualHost  string
	ExchangeName string
	ExchangeMode string
}

type AMQPOptions struct {
	// ConsumeMessageChanSize Set the channel size for consuming messages.
	ConsumeMessageChanSize int

	// ExchangeDurable Determines if the declared exchange will persist across server restarts (survives crashes).
	// true for durable, false for non-durable (transient).
	ExchangeDurable bool

	// ExchangeAutoDeleted Controls automatic deletion of the exchange when no longer in use.
	// true for auto-deletion, false for manual deletion.
	ExchangeAutoDeleted bool

	// ExchangeInternal Specifies if the exchange is accessible from outside the message broker (applicable to some systems).
	// true for internal, false for externally accessible.
	ExchangeInternal bool

	// ExchangeNoWait  Indicates whether the exchange declaration should be synchronous (waits for server response) or asynchronous.
	// true for no-wait (asynchronous), false for synchronous.
	ExchangeNoWait bool

	// ExchangeArgs An optional dictionary to provide exchange-specific configuration options.
	ExchangeArgs map[string]interface{}

	// QueueDurable Similar to ExchangeDurable, determines if the declared queue persists across server restarts.
	// true for durable, false for non-durable.
	QueueDurable bool

	// QueueAutoDeleted Similar to ExchangeAutoDeleted, controls automatic deletion of the queue.
	// true for auto-deletion, false for manual deletion.
	QueueAutoDeleted bool

	// QueueExclusive  Sets whether the queue should be created exclusively for the declaring connection.
	// true : The queue is created exclusively for the declaring connection. This means:
	// Only the connection that declared the queue can publish messages to it.
	// The queue will be automatically deleted when the declaring connection closes.
	// false: The queue is not exclusive and can be accessed by any connection. Multiple connections can publish and consume messages from the queue.
	QueueExclusive bool

	// QueueNoWait Similar to ExchangeNoWait, indicates synchronous or asynchronous queue declaration.
	// true for no-wait, false for synchronous.
	QueueNoWait bool

	// QueueArgs An optional dictionary to provide queue-specific configuration options
	QueueArgs map[string]interface{}

	// PublishMandatory Controls message delivery behavior on routing key mismatches.
	// true for mandatory publishing (raises an error if no queue matches),
	// false for non-mandatory (message silently discarded).
	PublishMandatory bool

	// PublishImmediate Determines if message publication should be confirmed immediately or asynchronously.
	// true for immediate confirmation (waits for server ack), false for asynchronous confirmation.
	PublishImmediate bool

	// ConsumerTag specifies a unique identifier for the consumer.
	// When you specify an empty string ("") for consumer_tag, the message queuing server typically assigns a unique identifier for that specific channel.
	// This means messages will be delivered to the consumer identified by the server-generated tag within that channel.
	ConsumerTag string

	// ConsumerAutoAck Controls automatic acknowledgment of delivered messages by the consumer.
	// true for auto-acknowledgment, false for manual acknowledgment (requires consumer to send ack messages).
	ConsumerAutoAck bool

	// ConsumerExclusive Indicates if the queue should be exclusive to this consumer.
	// true for exclusive access, false for shared access.
	ConsumerExclusive bool

	// ConsumerNoLocal Prevents the consumer from receiving messages published from the same connection.
	// true to avoid local messages, false to receive all messages.
	ConsumerNoLocal bool

	// ConsumerNoWait indicates synchronous or asynchronous consumer declaration.
	// true for no-wait, false for synchronous.
	ConsumerNoWait bool

	// ConsumerArgs An optional dictionary to provide consumer-specific configuration options
	ConsumerArgs map[string]interface{}
}

// NEWAMQPConfig generates the AMQPConfig from the AMQPConfig .
func NEWAMQPConfig(base AMQPBaseConfig, optional *AMQPOptions) AMQPConfig {
	AMQPcfg := AMQPConfig{}
	AMQPcfg.baseConfig = base
	if optional != nil {
		AMQPcfg.AMQPOption = *optional
	} else {
		AMQPcfg.AMQPOption.ExchangeDurable = true
		AMQPcfg.AMQPOption.QueueAutoDeleted = true
		AMQPcfg.AMQPOption.ConsumerAutoAck = true
		AMQPcfg.AMQPOption.ConsumerTag = ""
		AMQPcfg.AMQPOption.ConsumeMessageChanSize = 100
	}

	return AMQPcfg
}
