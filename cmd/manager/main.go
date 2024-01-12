package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
	usermock "github.com/ormushq/ormus/manager/mock"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/pkg/cryption"
)

func main() {
	cfg := config.C()

	setupSvc := setupServices(cfg)

	server := httpserver.New(cfg, setupSvc)

	server.Server()
}

func setupServices(cfg config.Config) httpserver.SetupServicesResponse {
	crypt := cryption.New(cryption.CryptConfing{})
	jwt := authservice.NewJWT(cfg.Manager.JWTConfig, crypt)
	unknownRepo := usermock.NewMockRepository(false)
	userSvc := userservice.New(jwt, unknownRepo)
	validateUserSvc := uservalidator.New(unknownRepo)

	userHand := userhandler.New(userSvc, validateUserSvc)

	return httpserver.SetupServicesResponse{
		UserHandler: userHand,
	}
}
