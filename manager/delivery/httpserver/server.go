package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
)

type SetupServicesResponse struct {
	UserHandler *userhandler.Handler
}

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
	Router      *echo.Echo
}

func New(cfg config.Config, setupSvc SetupServicesResponse) *Server {
	return &Server{
		config:      cfg,
		userHandler: *setupSvc.UserHandler,
		Router:      echo.New(),
	}
}

func (s *Server) Server() {
	e := echo.New()

	s.userHandler.SetUserRoute(e)

	e.GET("/health-check", s.healthCheck)

	e.Logger.Fatal(e.Start(":8080"))
}
