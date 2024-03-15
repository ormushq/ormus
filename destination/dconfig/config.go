package dconfig

type Config struct {
	RabbitMQTaskManagerConnection RabbitMQTaskManagerConnection `koanf:"rabbitmq_task_manager_connection"`
	RabbitMQConsumerConnection    RabbitMQConsumerConnection    `koanf:"rabbitmq_consumer_connection"`
	ConsumerTopic                 ConsumerTopic                 `koanf:"consumer_topic"`
	RedisTaskIdempotency          RedisTaskIdempotency          `koanf:"redis_idempotency"`
}
