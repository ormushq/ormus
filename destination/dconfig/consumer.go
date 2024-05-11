package dconfig

type ConsumerTopic string

type RabbitMQConsumerConnection struct {
	User            string `koanf:"user"`
	Password        string `koanf:"password"`
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	Vhost           string `koanf:"vhost"`
	ReconnectSecond int    `koanf:"reconnect_second"`
}
