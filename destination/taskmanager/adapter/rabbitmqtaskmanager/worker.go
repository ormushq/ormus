package rabbitmqtaskmanager

import (
	"fmt"
	"log"

	"github.com/ormushq/ormus/destination/entity"
	"github.com/ormushq/ormus/destination/integrationhandler"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Worker struct {
	TaskManager *TaskManager
	Handler     integrationhandler.IntegrationHandler
}

func NewWorker(tm *TaskManager, h integrationhandler.IntegrationHandler) *Worker {
	return &Worker{
		TaskManager: tm,
		Handler:     h,
	}
}

func (w *Worker) ProcessJobs() {
	connectionConfig := w.TaskManager.config

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", connectionConfig.User, connectionConfig.Password, connectionConfig.Host, connectionConfig.Port))
	panicOnWorkersError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	panicOnWorkersError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		w.TaskManager.queueName, // name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	panicOnWorkersError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	panicOnWorkersError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	panicOnWorkersError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {

			// todo should we ack message if we encounter any error ?
			// Acknowledge that message Received
			if err = d.Ack(false); err != nil {
				printWorkersError(err, "Failed to acknowledge message")
			}

			task, err := w.TaskManager.UnmarshalMessageToTask(d.Body)
			if err != nil {
				printWorkersError(err, "Failed to unmarshall message")
				continue
			}
			log.Printf("Task [%s] received by RabbitMQ worker.", task.ID)

			ti := w.TaskManager.taskIdempotency
			var taskStatus entity.TaskStatus
			taskID := task.ID

			enabled, err := ti.IntegrationHandlerIsEnable(taskID)
			if err != nil {
				log.Println("Error on IntegrationHandlerIsEnable.", err)
				continue
			}

			if enabled {
				err := w.Handler.Handle(task.ProcessedEvent)
				if err != nil {
					taskStatus = entity.FailedInIntegrationHandler
				}
			} else {
				log.Printf("\033[33mPrevent to handling duplicate processed event in idempotency.!\033[0m\n")
			}

			taskStatus = entity.SuccessInIntegrationHandler

			err = ti.Save(taskID, taskStatus)
			if err != nil {
				log.Println("Error on Saving task status.", err)

				continue
			}

		}
	}()

	log.Printf(" [RabbitMQ] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func panicOnWorkersError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func printWorkersError(err error, msg string) {
	log.Printf("%s: %s", msg, err)
}
