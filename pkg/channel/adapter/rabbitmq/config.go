package rbbitmqchannel

type RabbitMQConsumerConnection struct {
	User            string
	Password        string
	Host            string
	Port            int
	Vhost           string
	ReconnectSecond int
}
