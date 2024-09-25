package main

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/delivery/httpserver"
	"github.com/ormushq/ormus/manager/delivery/httpserver/projecthandler"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/ormushq/ormus/manager/repository/scyllarepo/scyllaproject"
	"github.com/ormushq/ormus/manager/repository/scyllarepo/scyllauser"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/projectvalidator"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/manager/workers"
	"github.com/ormushq/ormus/pkg/channel"
	"github.com/ormushq/ormus/pkg/channel/adapter/simple"
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
	logger.SetLevel(slog.LevelDebug)
	logger.L().Debug("start manger")
	cfg := config.C().Manager
	done := make(chan bool)
	wg := sync.WaitGroup{}
	logger.L().Debug(fmt.Sprintf("%+v", cfg))
	logger.L().Debug(fmt.Sprintf("%+v", cfg.ScyllaDBConfig))

	internalBroker := simple.New(done, &wg)
	err := internalBroker.NewChannel(managerparam.CreateDefaultProject, channel.BothMode,
		cfg.InternalBrokerConfig.ChannelSize, cfg.InternalBrokerConfig.NumberInstant, cfg.InternalBrokerConfig.MaxRetryPolicy)
	if err != nil {
		logger.L().Error("error on creating internal broker channel", err)
	}

	authSvc := authservice.New(cfg.AuthConfig)
	scylla, err := scyllarepo.New(cfg.ScyllaDBConfig)
	if err != nil {
		logger.L().Error("err msg:", err)
	}

	projectRepo := scyllaproject.New(scylla)
	projectValidator := projectvalidator.New(projectRepo)
	projectSvc := projectservice.New(projectRepo, internalBroker, projectValidator)
	projectHandler := projecthandler.New(authSvc, projectSvc)

	userRepo := scyllauser.New(scylla)
	userValidator := uservalidator.New(userRepo)
	userSvc := userservice.New(authSvc, userRepo, internalBroker, userValidator)
	userHand := userhandler.New(userSvc, projectSvc)

	workers.New(projectSvc, internalBroker).Run(done, &wg)

	server := httpserver.New(cfg, httpserver.SetupServicesResponse{
		UserHandler:    userHand,
		ProjectHandler: projectHandler,
	})

	server.Server()
}
