package dtcoordinator

import (
	"fmt"
	"log"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/destination/integrationhandler/adapters/fakeintegrationhandler"
	"github.com/ormushq/ormus/destination/taskmanager"
	"github.com/ormushq/ormus/destination/taskmanager/adapter/rabbitmqtaskmanager"
	"github.com/ormushq/ormus/destination/taskservice"
	"github.com/ormushq/ormus/event"
)

// DestinationTypeCoordinator is responsible for setup task managers and publish incoming processed events using suitable task publishers.

type DestinationTypeCoordinator struct {
	TaskService              taskservice.Service
	TaskPublishers           map[string]taskmanager.Publisher
	RabbitMQConnectionConfig dconfig.RabbitMQTaskManagerConnection
}

func New(ts taskservice.Service, rmqCnf dconfig.RabbitMQTaskManagerConnection) DestinationTypeCoordinator {
	// Create RabbitMQ task manager for webhook events
	rmqTaskManagerForWebhooks := rabbitmqtaskmanager.NewTaskManager(rmqCnf, "webhook_tasks_queue")

	taskPublishers := make(map[string]taskmanager.Publisher)
	taskPublishers["webhook"] = rmqTaskManagerForWebhooks

	return DestinationTypeCoordinator{
		TaskService:              ts,
		TaskPublishers:           taskPublishers,
		RabbitMQConnectionConfig: rmqCnf,
	}
}

func (d DestinationTypeCoordinator) Start(processedEvents <-chan event.ProcessedEvent, done <-chan bool, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Start task coordinator [DestinationType].")

		for {
			select {
			case pe := <-processedEvents:

				taskPublisher, ok := d.TaskPublishers["webhook"]
				if !ok {
					slog.Error(fmt.Sprintf("Error on finding task manager for %s", pe.DestinationType()))

					break
				}

				pErr := taskPublisher.Publish(pe)
				if pErr != nil {
					fmt.Println(pErr)

					break
				}

			case <-done:

				return
			}
		}
	}()

	webhookTaskConsumer := rabbitmqtaskmanager.NewTaskConsumer(d.RabbitMQConnectionConfig, "webhook_tasks_queue")
	fakeWebhookHandler := fakeintegrationhandler.New()

	// Run workers
	// todo we can use loop in range of slices of workers.
	// also we can use config for number of each worker for different destination types.

	webhookWorker1 := rabbitmqtaskmanager.NewWorker(webhookTaskConsumer, fakeWebhookHandler, d.TaskService)
	err := webhookWorker1.Run(done, wg)
	if err != nil {
		log.Panicf("%s: %s", "Error on webhook worker", err)

		return err
	}

	return nil
}
