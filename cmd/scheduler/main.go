package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/rabbitmqchannel"
	"github.com/ormushq/ormus/scheduler"
	sourceevent "github.com/ormushq/ormus/source/eventhandler"
	"github.com/ormushq/ormus/source/repository/scylladb"
	schrepo "github.com/ormushq/ormus/source/repository/scylladb/scheduler"
	sourcescheduler "github.com/ormushq/ormus/source/scheduler"
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
	ctx, cancel := context.WithCancel(context.Background())
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

	cfg := config.C()
	sourceSch := SetupSourceServices(cfg, done, wg)
	sch := scheduler.New(done, wg)
	wg.Add(1)
	go func() {
		sch.Start(ctx, cfg.Source.UndeliveredEventRetransmitPeriod, sourceSch.PublishUndeliveredEvents)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	logger.L().Info("Received shutdown signal, shutting down gracefully...")
	cancel()
	<-done
	logger.L().Info("system scheduler successfully shut down gracefully")
	wg.Wait()
}

func SetupSourceServices(cfg config.Config, done chan bool, wg *sync.WaitGroup) sourcescheduler.Scheduler {
	inputAdapter := rabbitmqchannel.New(done, wg, cfg.RabbitMq)
	err := inputAdapter.NewChannel(cfg.Source.NewEventQueueName, channel.InputOnlyMode, cfg.Source.BufferSize, cfg.Source.MaxRetry)
	if err != nil {
		panic(err)
	}

	Publisher := sourceevent.NewPublisher(inputAdapter)

	DB, err := scylladb.New(cfg.Source.ScyllaDBConfig)
	if err != nil {
		panic(err)
	}

	schRepo := schrepo.New(DB)
	schSvc := sourcescheduler.New(schRepo, Publisher, cfg.Source, wg)

	return schSvc
}
