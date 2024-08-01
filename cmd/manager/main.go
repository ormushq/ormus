package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/delivery/grpcserver"
	"github.com/ormushq/ormus/manager/delivery/httpserver"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/mockRepo/projectstub"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/ormushq/ormus/manager/repository/sourcerepo"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/manager/workers"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/simple"
)

func main() {
	logger.SetLevel(slog.LevelDebug)
	logger.L().Debug("start manger")
	cfg := config.C().Manager
	done := make(chan bool)
	wg := sync.WaitGroup{}
	fmt.Println(cfg.ScyllaDBConfig)

	internalBroker := simple.New(done, &wg)
	err := internalBroker.NewChannel(managerparam.CreateDefaultProject, channel.BothMode,
		cfg.InternalBrokerConfig.ChannelSize, cfg.InternalBrokerConfig.NumberInstant, cfg.InternalBrokerConfig.MaxRetryPolicy)
	if err != nil {
		logger.L().Error("error on creating internal broker channel", err)
	}

	jwt := authservice.NewJWT(cfg.JWTConfig)
	scylla, err := scyllarepo.New(cfg.ScyllaDBConfig)
	if err != nil {
		logger.L().Error("err msg:", err)
	}

	unknownRepo1 := projectstub.New(false)

	ProjectSvc := projectservice.New(&unknownRepo1, internalBroker)

	userSvc := userservice.New(jwt, scylla, internalBroker)

	validateUserSvc := uservalidator.New(scylla)

	userHand := userhandler.New(userSvc, validateUserSvc, ProjectSvc)
	workers.New(ProjectSvc, internalBroker).Run(done, &wg)

	httpServer := httpserver.New(cfg, httpserver.SetupServicesResponse{
		UserHandler: userHand,
	})

	sourceRepo := sourcerepo.New(scylla)

	grpcServer := grpcserver.New(sourceRepo, cfg)

	if err := httpServer.Start(); err != nil {
		logger.L().Error("Failed to start manager HTTP server", err)
		os.Exit(1)
	}

	if err := grpcServer.Start(); err != nil {
		logger.L().Error("Failed to start gRPC server", "error", err)
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.L().Info("Shutting down manager server...")

	const timeoutDuration = 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.L().Error("manager HTTP server forced to shutdown", "error", err)
	}

	if err := grpcServer.Stop(ctx); err != nil {
		logger.L().Error("manager gRPC server forced to shutdown", "error", err)
	}

	logger.L().Info("Server exiting")
}
