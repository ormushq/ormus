package destination

import "github.com/ormushq/ormus/destination/config"

type Config struct {
	RabbitMQTaskManagerConnection config.RabbitMQTaskManagerConnection `koanf:"rabbitmq_task_manager_connection"`
	RabbitMQConsumerConnection    config.RabbitMQConsumerConnection    `koanf:"rabbitmq_consumer_connection"`
	ConsumerTopic                 config.ConsumerTopic                 `koanf:"consumer_topic"`
}
