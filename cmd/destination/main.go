package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/processedevent/adapter/rabbitmqconsumer"
	"github.com/ormushq/ormus/destination/taskcoordinator/adapter/dtcoordinator"
	"github.com/ormushq/ormus/destination/taskservice"
	"github.com/ormushq/ormus/destination/taskservice/adapter/idempotency/redistaskidempotency"
	"github.com/ormushq/ormus/destination/taskservice/adapter/repository/inmemorytaskrepo"
	"github.com/ormushq/ormus/logger"
)

const waitingAfterShutdownInSeconds = 2

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	fileMaxSizeInMB := 10
	fileMaxAgeInDays := 30

	//------ Setup logger ------
	cfg := logger.Config{
		FilePath:         "./destination/logs.json",
		UseLocalTime:     false,
		FileMaxSizeInMB:  fileMaxSizeInMB,
		FileMaxAgeInDays: fileMaxAgeInDays,
	}

	logLevel := slog.LevelInfo
	if config.C().Destination.DebugMode {
		logLevel = slog.LevelDebug
	}

	opt := slog.HandlerOptions{
		// todo should level debug be read from config?
		Level: logLevel,
	}
	l := logger.New(cfg, &opt)
	slog.SetDefault(l)

	// Setup Task Service

	// In-Memory task idempotency
	// taskIdempotency := inmemorytaskidempotency.New()

	// Redis task idempotency
	// todo do we need to use separate db number for redis task idempotency or destination module?
	redisAdapter, err := redis.New(config.C().Redis)
	if err != nil {
		log.Panicf("error in new redis")
	}
	taskIdempotency := redistaskidempotency.New(redisAdapter, "tasks:", 30*24*time.Hour)
	taskRepo := inmemorytaskrepo.New()
	taskService := taskservice.New(taskIdempotency, taskRepo)
	//----- Consuming processed events -----//

	// Get connection config for rabbitMQ consumer
	rmqConsumerConnConfig := config.C().Destination.RabbitMQConsumerConnection
	rmqConsumerTopic := config.C().Destination.ConsumerTopic

	// todo should we consider array of topics?
	rmqConsumer := rabbitmqconsumer.New(rmqConsumerConnConfig, rmqConsumerTopic)

	log.Println("Start Consuming processed events.")
	processedEvents, err := rmqConsumer.Consume(done, &wg)
	if err != nil {
		log.Panicf("Error on consuming processed events.")
	}

	//----- Setup Task Coordinator -----//
	// Task coordinator specifies which task manager should handle incoming processed events.
	// we can have different task coordinators base on destination type, customer plans, etc.
	// Now we just create dtcoordinator that stands for destination type coordinator.
	// It determines which task manager should be used for processed evens considering destination type of processed events.

	// todo maybe it is better to having configs for setup of task coordinator.

	rmqTaskManagerConnConfig := config.C().Destination.RabbitMQTaskManagerConnection
	coordinator := dtcoordinator.New(taskService, rmqTaskManagerConnConfig)

	cErr := coordinator.Start(processedEvents, done, &wg)
	if cErr != nil {
		log.Panicf("Error on starting destination type coordinator.")
	}

	//----- Handling graceful shutdown  -----//

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Received interrupt signal, shutting down gracefully...")
	done <- true

	close(done)

	// todo use config for waiting time after graceful shutdown
	time.Sleep(waitingAfterShutdownInSeconds * time.Second)
	wg.Wait()
}
