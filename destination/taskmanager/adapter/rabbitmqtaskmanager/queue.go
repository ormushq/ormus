package rabbitmqtaskmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ormushq/ormus/destination/config"
	"github.com/ormushq/ormus/destination/entity"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Queue struct {
	name   string
	config config.RabbitMQTaskManagerConnection
}

func NewQueue(c config.RabbitMQTaskManagerConnection, n string) *Queue {
	return &Queue{
		name:   n,
		config: c,
	}
}

func (q *Queue) Enqueue(task *entity.Task) error {

	connectionConfig := q.config
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", connectionConfig.User, connectionConfig.Password, connectionConfig.Host, connectionConfig.Port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	rq, err := ch.QueueDeclare(
		q.name, // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonTask, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = ch.PublishWithContext(ctx,
		"",      // exchange
		rq.Name, // routing key
		false,   // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         jsonTask,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf("Task [%s] is published to RabbitMQ queue.", task.ID)

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
