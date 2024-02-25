package rabbitmqconsumer

import (
	"encoding/json"
	"fmt"
	gc "github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/config"
	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/destination/taskmanager/adapter/rabbitmqtaskmanager"
	"github.com/ormushq/ormus/destination/taskstorage/adapter/redistaskstorage"
	"github.com/ormushq/ormus/event"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Consumer struct {
	topic            config.ConsumerTopic
	connectionConfig config.RabbitMQConsumerConnection
}

func NewConsumer(t config.ConsumerTopic, c config.RabbitMQConsumerConnection) *Consumer {
	return &Consumer{
		connectionConfig: c,
		topic:            t,
	}
}

func (c *Consumer) Consume() error {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", c.connectionConfig.User, c.connectionConfig.Password, c.connectionConfig.Host, c.connectionConfig.Port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"processed_events_exchange", // name
		"topic",                     // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Binding queue %s to exchange", q.Name)
	err = ch.QueueBind(
		q.Name,                      // queue name
		string(c.topic),             // routing key
		"processed_events_exchange", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var pe event.ProcessedEvent
			if err := json.Unmarshal(d.Body, &pe); err != nil {
				fmt.Println("Error on unMarshaling processed event:", err)
			}
			//TODO how to get corresponding task manager? I think we need strategy of factory design pattern or something like pool of task managers.
			// we need to find suitable task manager using consumer_topic or even processed event info.
			// but now for simplicity I used hard coded approach.

			ts := redistaskstorage.New()
			rmqTaskManagerConnConfig := gc.C().Destination.RabbitMQTaskManagerConnection
			rmqTaskManager := rabbitmqtaskmanager.NewTaskManager(rmqTaskManagerConnConfig, "webhook_queue", ts)
			webhookTask := GenerateTaskUsingProcessedEvent(pe)
			err := rmqTaskManager.SendToQueue(&webhookTask)
			if err != nil {
				log.Println("Error on send task to queue. ", err)
			}

			log.Println("Processed event received by RabbitMQ consumer.")
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GenerateTaskUsingProcessedEvent(pe event.ProcessedEvent) entity.Task {
	return entity.Task{
		ID:             pe.MessageID + "-" + pe.Integration.ID,
		Name:           pe.Integration.Metadata.Slug + "_integration_delivery", // ex : webhook_integration_delivery
		ProcessedEvent: pe,
		Timestamp:      time.Time{},
	}
}
