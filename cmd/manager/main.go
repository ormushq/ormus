package main

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/delivery/httpserver"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/mockRepo/projectstub"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
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

	validateUserSvc := uservalidator.New(scylla)

	userSvc := userservice.New(jwt, scylla, internalBroker, validateUserSvc)

	userHand := userhandler.New(userSvc, ProjectSvc)

	workers.New(ProjectSvc, internalBroker).Run(done, &wg)

	server := httpserver.New(cfg, httpserver.SetupServicesResponse{
		UserHandler: userHand,
	})

	server.Server()
}
