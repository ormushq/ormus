package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Adapter struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	messages   chan *Message
}
type Config struct {
	UserName string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
}

type Message struct {
	Body        []byte
	DeliveryTag uint64 // Delivery tag for acknowledgment
}

func New(cfg Config) (Adapter, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.UserName, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return Adapter{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Adapter{}, err
	}

	return Adapter{
		connection: conn,
		channel:    ch,
		messages:   make(chan *Message), // Initialize channel
	}, nil

}

func (a Adapter) GetChannel() *amqp.Channel {
	return a.channel

}

func (a Adapter) GetMessage() chan *Message {
	return a.messages

}

func (a Adapter) GetConnection() *amqp.Connection {
	return a.connection
	
}
