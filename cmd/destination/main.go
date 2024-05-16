package main

import (
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/processedevent/adapter/rabbitmqconsumer"
	"github.com/ormushq/ormus/destination/taskcoordinator/adapter/dtcoordinator"
	"github.com/ormushq/ormus/destination/taskmanager/adapter/rabbitmqchanneltaskmanager"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
)

const waitingAfterShutdownInSeconds = 1

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	//----------------- Setup Logger -----------------//

	fileMaxSizeInMB := 10
	fileMaxAgeInDays := 30

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

	// use slog as default logger.
	slog.SetDefault(l)

	//----------------- Consume Processed Events From Core -----------------//

	// Get connection config for rabbitMQ consumer
	rmqConsumerConnConfig := config.C().Destination.RabbitMQConsumerConnection
	rmqConsumerTopic := config.C().Destination.ConsumerTopic

	// todo should we consider array of topics?
	rmqConsumer := rabbitmqconsumer.New(rmqConsumerConnConfig, rmqConsumerTopic)

	slog.Info("Start Consuming processed events.")
	processedEvents, err := rmqConsumer.Consume(done, &wg)
	if err != nil {
		log.Panicf("Error on consuming processed events.")
	}

	//----------------- Setup Task Coordinator -----------------//

	// Task coordinator is responsible for considering task's characteristics
	// and publish it to task queue using task publisher. currently we support
	// destination type coordinator which means every task with specific destination type
	// will be published to its corresponding task publisher.

	reconnectSecond := 5
	channelSize := 100
	numberInstant := 5
	maxRetryPolicy := 5

	taskPublisherCnf := config.C().Destination.RabbitMQTaskManagerConnection

	inputChannelAdapter := rbbitmqchannel.New(done, &wg, dconfig.RabbitMQConsumerConnection{
		User:            taskPublisherCnf.User,
		Password:        taskPublisherCnf.Password,
		Host:            taskPublisherCnf.Host,
		Port:            taskPublisherCnf.Port,
		Vhost:           "/",
		ReconnectSecond: reconnectSecond,
	})

	webHookQueueName := "webhook_tasks"

	inputChannelAdapter.NewChannel(webHookQueueName, channel.InputOnlyMode, channelSize, numberInstant, maxRetryPolicy)

	inputChannel, err := inputChannelAdapter.GetInputChannel(webHookQueueName)
	if err != nil {
		log.Fatalf("Couldn't get input channel for %s: %s", webHookQueueName, err)
	}

	webhookTaskPublisher := rabbitmqchanneltaskmanager.NewTaskPublisher(inputChannel)

	taskPublishers := make(dtcoordinator.TaskPublisherMap)
	taskPublishers[entity.WebhookDestinationType] = webhookTaskPublisher

	coordinator := dtcoordinator.New(taskPublishers)

	cErr := coordinator.Start(processedEvents, done, &wg)
	if cErr != nil {
		log.Panicf("Error on starting destination type coordinator.")
	}

	//----------------- Handling graceful shutdown -----------------//

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	slog.Info("Received interrupt signal, shutting down gracefully...")
	done <- true

	close(done)

	// todo use config for waiting time after graceful shutdown
	time.Sleep(waitingAfterShutdownInSeconds * time.Second)
	wg.Wait()
}
