package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/processedevent/adapter/rabbitmqconsumer"
	"github.com/ormushq/ormus/destination/taskcoordinator/adapter/dtcoordinator"
	"github.com/ormushq/ormus/destination/taskservice"
	"github.com/ormushq/ormus/destination/taskservice/adapter/idempotency/inmemorytaskidempotency"
	"github.com/ormushq/ormus/destination/taskservice/adapter/repository/inmemorytaskrepo"
)

const waitingAfterShutdownInSeconds = 2

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	// Setup Task Service
	taskIdempotency := inmemorytaskidempotency.New()
	taskRepo := inmemorytaskrepo.New()
	taskService := taskservice.New(taskIdempotency, taskRepo)

	//----- Consume processed events -----//

	// Get connection config for rabbitMQ consumer
	rmqConsumerConnConfig := config.C().Destination.RabbitMQConsumerConnection
	rmqConsumerTopic := config.C().Destination.ConsumerTopic

	rmqConsumer := rabbitmqconsumer.NewConsumer(rmqConsumerConnConfig, rmqConsumerTopic)

	log.Println("Consuming processed events...")
	processedEvents, err := rmqConsumer.Consume(done, &wg)
	if err != nil {
		log.Panicf("Error on consuming processed events")
	}

	//----- Setup Task Coordinator-----//

	// todo we can use configs to create different dispatcher base one destination type or customer plans and ...

	// dtcoordinator stands for destination type coordinator.
	// It determines which task manager should be used for processed evens considering destination type.
	rmqTaskManagerConnConfig := config.C().Destination.RabbitMQTaskManagerConnection
	coordinator := dtcoordinator.New(taskService, rmqTaskManagerConnConfig)

	cErr := coordinator.Start(processedEvents, done, &wg)
	if cErr != nil {
		log.Panicf("Error on starting destination type coordinator")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Received interrupt signal, shutting down gracefully..")
	done <- true

	close(done)

	// todo use config for waiting time after graceful shutdown
	time.Sleep(waitingAfterShutdownInSeconds * time.Second)
	wg.Wait()
}
