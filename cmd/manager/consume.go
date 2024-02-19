package main

import (
	"fmt"
	"github.com/ormushq/ormus/pkg/broker/rabbitmq"
)

func main() {
	fmt.Println("consumer app")
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
	chmsg, err := conn.ConsumeMessage("sina")
	if err != nil {
		panic(err)
	}
	for ts := range chmsg {
		fmt.Println("ID:", ts.ID, "topic:", ts.Topic, "payload:", string(ts.Payload))
	}

}
