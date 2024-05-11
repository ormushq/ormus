package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/etcd"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/taskdelivery"
	"github.com/ormushq/ormus/destination/taskdelivery/adapters/fakedeliveryhandler"
	"github.com/ormushq/ormus/destination/taskmanager/adapter/rabbitmqtaskmanager"
	"github.com/ormushq/ormus/destination/taskservice"
	"github.com/ormushq/ormus/destination/taskservice/adapter/idempotency/redistaskidempotency"
	"github.com/ormushq/ormus/destination/taskservice/adapter/repository/inmemorytaskrepo"
	"github.com/ormushq/ormus/destination/worker"
	"github.com/ormushq/ormus/logger"
)

const waitingAfterShutdownInSeconds = 2

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	//----------------- Setup Logger -----------------//

	fileMaxSizeInMB := 10
	fileMaxAgeInDays := 30

	cfg := logger.Config{
		FilePath:         "./destination/worker_log.json",
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

	//----------------- Setup Task Service -----------------//

	redisAdapter, err := redis.New(config.C().Redis)
	if err != nil {
		log.Panicf("error in new redis")
	}
	taskIdempotency := redistaskidempotency.New(redisAdapter, "tasks:", 30*24*time.Hour)

	taskRepo := inmemorytaskrepo.New()

	// Set up etcd as distributed locker.
	distributedLocker, err := etcd.New(config.C().Etcd)
	if err != nil {
		log.Panicf("Error on new etcd")
	}

	taskHandler := taskservice.New(taskIdempotency, taskRepo, distributedLocker)

	// Register delivery handlers
	// each destination type can have specific delivery handler
	fakeTaskDeliveryHandler := fakedeliveryhandler.New()
	taskdelivery.Register("webhook", fakeTaskDeliveryHandler)

	//----------------- Consume ProcessEvents -----------------//

	taskConsumerConf := config.C().Destination.RabbitMQTaskManagerConnection
	webhookTaskConsumer := rabbitmqtaskmanager.NewTaskConsumer(taskConsumerConf, "webhook_tasks_queue")

	processedEvents, err := webhookTaskConsumer.Consume(done, &wg)
	if err != nil {
		log.Panicf("Error on consuming tasks.")
	}

	w1 := worker.NewWorker(processedEvents, taskHandler)

	w1Err := w1.Run(done, &wg)
	if w1Err != nil {
		log.Panicf("%s: %s", "Error on webhook worker", err)
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
