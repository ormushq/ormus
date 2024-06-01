package main

import (
	"sync"
	"time"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/mockRepo/projectstub"
	"github.com/ormushq/ormus/manager/mockRepo/usermock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/manager/workers"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/simple"
)

func main() {
	cfg := config.C().Manager
	done := make(chan bool)
	wg := sync.WaitGroup{}

	internalBroker := simple.New(done, &wg)
	internalBroker.NewChannel("CreateDefaultProject", channel.BothMode,
		cfg.InternalBrokerConfig.ChannelSize, cfg.InternalBrokerConfig.NumberInstant, cfg.InternalBrokerConfig.MaxRetryPolicy)

	TimeOfDayByHour := 24
	cfg.JWTConfig.AccessExpirationTimeInDay *= time.Duration(TimeOfDayByHour * int(time.Hour))
	cfg.JWTConfig.RefreshExpirationTimeInDay *= time.Duration(TimeOfDayByHour * int(time.Hour))

	jwt := authservice.NewJWT(cfg.JWTConfig)

	unknownRepo := usermock.NewMockRepository(false)
	unknownRepo1 := projectstub.New(false)

	ProjectSvc := projectservice.New(&unknownRepo1, internalBroker)

	userSvc := userservice.New(jwt, unknownRepo, internalBroker)

	validateUserSvc := uservalidator.New(unknownRepo)

	userHand := userhandler.New(userSvc, validateUserSvc, ProjectSvc)
	workers.New(ProjectSvc, internalBroker).Run(done, &wg)

	server := httpserver.New(cfg, httpserver.SetupServicesResponse{
		UserHandler: userHand,
	})

	server.Server()
}
