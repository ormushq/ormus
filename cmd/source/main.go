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
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	rbbitmqchannel "github.com/ormushq/ormus/pkg/channel/adapter/rabbitmq"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/delivery/httpserver/statushandler"
	sourceevent "github.com/ormushq/ormus/source/eventhandler"
	writekeyrepo "github.com/ormushq/ormus/source/repository/redis/rediswritekey"
	"github.com/ormushq/ormus/source/service/writekey"
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

	handlers := []httpserver.Handler{
		statushandler.New(),
	}
	err := otela.Configure(wg, done, otela.Config{Exporter: otela.ExporterConsole})
	if err != nil {
		panic(err.Error())
	}

	cfg := config.C()
	_, Consumer := SetupSourceServices(cfg)
	Consumer.Consume(context.Background(), cfg.Source.NewSourceEventName, done, wg, Consumer.ProcessNewSourceEvent)

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

func SetupSourceServices(cfg config.Config) (writekey.Service, sourceevent.Consumer) {
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	dbConfig := dconfig.RabbitMQConsumerConnection{
		User:            cfg.RabbitMq.UserName,
		Password:        cfg.RabbitMq.Password,
		Host:            cfg.RabbitMq.Host,
		Port:            cfg.RabbitMq.Port,
		Vhost:           cfg.RabbitMq.Vhost,
		ReconnectSecond: cfg.RabbitMq.ReconnectSecond,
	}
	outputAdapter := rbbitmqchannel.New(done, wg, dbConfig)
	err := outputAdapter.NewChannel(cfg.Source.NewSourceEventName, channel.OutputOnly, cfg.Source.BufferSize, cfg.Source.NumberInstants, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}

	adapter, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	writeKeyRepo := writekeyrepo.New(adapter)
	writeKeySvc := writekey.New(&writeKeyRepo, cfg.Source)
	eventHandler := sourceevent.New(outputAdapter, writeKeySvc)

	return writeKeySvc, *eventHandler
}
