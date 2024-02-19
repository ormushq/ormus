package main

import (
	"fmt"
	"github.com/ormushq/ormus/pkg/broker/message_broker"
	"github.com/ormushq/ormus/pkg/broker/rabbitmq"
)

func main() {

	amqpCgq := rabbitmq.DefaultAMQPConfig()
	conn, err := rabbitmq.NewRabbitMQBroker(amqpCgq)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("connect")
	err = conn.DeclareExchange("test1", "direct")
	if err != nil {
		panic(err)
	}
	err = conn.PublishMessage("sina", message_broker.NewMessage("tet", []byte("hello world")))
	if err != nil {
		panic(err)
	}
	fmt.Println("message send...")
}
