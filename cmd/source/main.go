package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/rabbitmqchannel"
	"github.com/ormushq/ormus/source/adapter/manager"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/delivery/httpserver/eventhandler"
	"github.com/ormushq/ormus/source/delivery/httpserver/statushandler"
	sourceevent "github.com/ormushq/ormus/source/eventhandler"
	writekeyrepo "github.com/ormushq/ormus/source/repository/redis/rediswritekey"
	"github.com/ormushq/ormus/source/repository/scylladb"
	eventrepo "github.com/ormushq/ormus/source/repository/scylladb/event"
	eventsvc "github.com/ormushq/ormus/source/service/event"
	"github.com/ormushq/ormus/source/service/writekey"
	"github.com/ormushq/ormus/source/validator/eventvalidator/eventvalidator"
)

//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	JWTToken
//	@in							header
//	@name						Authorization

func main() {
	done := make(chan bool)
	wg := &sync.WaitGroup{}

	//----------------- Setup Logger -----------------//
	fileMaxSizeInMB := 10
	fileMaxAgeInDays := 30

	loggercfg := logger.Config{
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
	l := logger.New(loggercfg, &opt)

	// use slog as default logger.
	slog.SetDefault(l)

	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}

	cfg := config.C()
	_, consumer, eventSvc, eventValidator := SetupSourceServices(cfg)
	consumer.Consume(context.Background(), cfg.Source.NewSourceEventName, done, wg, consumer.ProcessNewSourceEvent)
	consumer.Consume(context.Background(), cfg.Source.UndeliveredEventsQueueName, done, wg, consumer.EventHasDeliveredToDestination)
	//----------------- Setup Tracer -----------------//
	otelcfg := otela.Config{
		Endpoint:           config.C().Source.Otel.Endpoint,
		ServiceName:        config.C().Source.Otel.ServiceName,
		EnableMetricExpose: config.C().Source.Otel.EnableMetricExpose,
		MetricExposePath:   config.C().Source.Otel.MetricExposePath,
		MetricExposePort:   config.C().Source.Otel.MetricExposePort,
		Exporter:           otela.ExporterGrpc,
	}
	err = otela.Configure(wg, done, otelcfg)
	if err != nil {
		l.Error(err.Error())
	}
	handlers := []httpserver.Handler{
		statushandler.New(),
		eventhandler.New(eventSvc, eventValidator),
	}
	httpServer := httpserver.New(config.C().Source, handlers)

	httpServer.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received interrupt signal, shutting down gracefully...")
	done <- true

	close(done)
	wg.Wait()
}

func SetupSourceServices(cfg config.Config) (writeKeySvc writekey.Service, consumer sourceevent.Consumer, eventSvc eventsvc.Service, eventValidator eventvalidator.Validator) {
	done := make(chan bool)
	wg := &sync.WaitGroup{}

	outputAdapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err := outputAdapter.NewChannel(cfg.Source.NewSourceEventName, channel.OutputOnly, cfg.Source.BufferSize, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}

	err = outputAdapter.NewChannel(cfg.Source.UndeliveredEventsQueueName, channel.OutputOnly, cfg.Source.BufferSize, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}

	inputAdapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err = inputAdapter.NewChannel(cfg.Source.NewEventQueueName, channel.InputOnlyMode, cfg.Source.BufferSize, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}

	Publisher := sourceevent.NewPublisher(inputAdapter)

	redisAdapter, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	ManagerAdapter := manager.New(cfg.Source)

	writeKeyRepo := writekeyrepo.New(redisAdapter, *ManagerAdapter)
	writeKeySvc = writekey.New(&writeKeyRepo, cfg.Source)

	DB, err := scylladb.New(cfg.Source.ScyllaDBConfig)
	if err != nil {
		panic(err)
	}
	eventRepo := eventrepo.New(DB, Publisher)
	eventSvc = *eventsvc.New(eventRepo, cfg.Source, wg)

	eventValidator = eventvalidator.New(&writeKeyRepo, cfg.Source)

	consumer = *sourceevent.NewConsumer(outputAdapter, writeKeySvc, eventSvc, cfg.Source.RetryNumber)

	return writeKeySvc, consumer, eventSvc, eventValidator
}
