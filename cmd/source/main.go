package main

import (
	"context"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/adapter/redis"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/adapter/rabbitmq"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/delivery/httpserver/statushandler"
	sourceevent "github.com/ormushq/ormus/source/eventhandler"
	writekeyrepo "github.com/ormushq/ormus/source/repository/redis/rediswritekey"
	"github.com/ormushq/ormus/source/service/writekey"
	"log/slog"
	"os"
	"os/signal"
	"sync"
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

	cfg := config.Config{
		Source: source.Config{
			AMPQURI:                 "amqp://guest:guest@localhost:5672/",
			WriteKeyRedisExpiration: 60,
		},
		Redis: redis.Config{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}

	_, Consumer := SetupSourceServices(cfg)
	Consumer.ConsumeWriteKey(context.Background(), "new-source-event", done, wg)

	//----------------- Setup Tracer -----------------//
	otelcfg := otela.Config{
		Endpoint:           config.C().Source.Otel.Endpoint,
		ServiceName:        config.C().Source.Otel.ServiceName,
		EnableMetricExpose: config.C().Source.Otel.EnableMetricExpose,
		MetricExposePath:   config.C().Source.Otel.MetricExposePath,
		MetricExposePort:   config.C().Source.Otel.MetricExposePort,
		Exporter:           otela.ExporterGrpc,
	}
	err := otela.Configure(wg, done, otelcfg)
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
	pub, err := rabbitmq.NewRabbitMQAdapter(cfg.Source.AMPQURI)
	if err != nil {
		panic(err)
	}
	sub, err := rabbitmq.NewRabbitMQAdapter(cfg.Source.AMPQURI)
	if err != nil {
		panic(err)
	}
	adapter, err := redis.New(cfg.Redis)
	if err != nil {
		panic(err)
	}

	writekeyrepo := writekeyrepo.New(adapter)
	writekeysvc := writekey.New(pub, sub, &writekeyrepo, cfg.Source)
	writekeysvc.CreateNewWriteKey(context.Background(), "1", "20", "3")
	eventhandler := sourceevent.New(sub, writekeysvc)

	return writekeysvc, *eventhandler
}
