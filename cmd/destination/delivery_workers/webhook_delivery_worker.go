package main

import (
	"context"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/destination/taskservice/adapter/idempotency/redistaskidempotency"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/destination/taskdelivery/adapters/webhookdeliveryhandler"
	"github.com/ormushq/ormus/destination/taskmanager/adapter/rabbitmqchanneltaskmanager"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"

	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ormushq/ormus/adapter/etcd"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/destination/taskdelivery"
	"github.com/ormushq/ormus/destination/taskservice"
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

	//----------------- Setup Tracer -----------------//
	otelcfg := otela.Config{
		Endpoint:           config.C().Destination.Otel.Endpoint,
		ServiceName:        config.C().Destination.Otel.ServiceName + "/WebhookDeliveryWorker",
		EnableMetricExpose: false,
		Exporter:           otela.ExporterGrpc,
	}

	err := otela.Configure(&wg, done, otelcfg)
	if err != nil {
		l.Error(err.Error())
	}

	tracer := otela.NewTracer("main")
	ctx, span := tracer.Start(context.Background(), "webhook-delivery-worker-main")

	//----------------- Setup Task Service -----------------//

	redisAdapter, err := redis.New(config.C().Redis)
	if err != nil {
		log.Panicf("error in new redis")
	}
	taskIdempotency := redistaskidempotency.New(redisAdapter, "tasks:", 30*24*time.Hour)
	span.AddEvent("task-idempotency-created")

	taskRepo := inmemorytaskrepo.New()
	span.AddEvent("task-repo-created")

	// Set up etcd as distributed locker.
	distributedLocker, err := etcd.NewWithContext(ctx, config.C().Etcd)
	if err != nil {
		log.Panicf("Error on new etcd")
	}
	span.AddEvent("etcd-created")

	taskHandler := taskservice.New(taskIdempotency, taskRepo, distributedLocker)
	span.AddEvent("task-handler-created")

	// Register delivery handlers
	// each destination type can have specific delivery handler
	webhookTaskDeliveryHandler := webhookdeliveryhandler.New()

	span.AddEvent("webhook-task-delivery-handler-created")

	taskdelivery.Register("webhook", webhookTaskDeliveryHandler)
	span.AddEvent("faker-task-delivery-handler-registered")

	//----------------- Consume ProcessEvents -----------------//

	channelSize := 100
	reconnectSecond := 10
	numberInstant := 5
	maxRetryPolicy := 5
	taskConsumerConf := config.C().Destination.RabbitMQTaskManagerConnection
	queueName := "webhook_tasks"
	outputChannelAdapter := rbbitmqchannel.NewWithContext(ctx, done, &wg, dconfig.RabbitMQConsumerConnection{
		User:            taskConsumerConf.User,
		Password:        taskConsumerConf.Password,
		Host:            taskConsumerConf.Host,
		Port:            taskConsumerConf.Port,
		Vhost:           "/",
		ReconnectSecond: reconnectSecond,
	})
	span.AddEvent("output-channel-adapter-created")

	errNCA := outputChannelAdapter.NewChannelWithContext(ctx, queueName, channel.OutputOnly, channelSize, numberInstant, maxRetryPolicy)
	if errNCA != nil {
		logger.WithGroup("webhook_delivery_worker").Error(errNCA.Error(),
			slog.String("error", errNCA.Error()))
		span.AddEvent("error-on-new-channel", trace.WithAttributes(
			attribute.String("error", errNCA.Error()),
		))
		os.Exit(1)
	}
	span.AddEvent("output-channel-created")

	outputChannel, err := outputChannelAdapter.GetOutputChannel(queueName)
	if err != nil {
		span.AddEvent("error-on-get-output-channel", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
		log.Panicf("Error on get output channel: %s", err)
	}

	webhookTaskConsumer := rabbitmqchanneltaskmanager.NewTaskConsumer(outputChannel, channelSize)
	span.AddEvent("webhookTaskConsumer-created")

	processedEvents, err := webhookTaskConsumer.Consume(done, &wg)
	if err != nil {
		log.Panicf("Error on consuming tasks.")
	}
	span.AddEvent("processedEvents-channel-opened")

	w1 := worker.NewWorker(processedEvents, taskHandler)
	span.AddEvent("worker-created")

	w1Err := w1.Run(done, &wg)
	if w1Err != nil {
		log.Panicf("%s: %s", "Error on webhook worker", err)
	}
	span.AddEvent("worker-started")

	span.End()

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
