package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/config"
	"github.com/ormushq/ormus/source/delivery/httpserver/userhandler"
)

type Server struct {
	config      config.Config
	Router      *echo.Echo
	userhandler userhandler.Handler
}

func New(c config.Config) Server {
	return Server{
		config: config.Config{
			HTTPServer: c.HTTPServer,
		},
		Router: echo.New(),
	}
}

func (s Server) Serve() {
	s.userhandler.SetRoutes(s.Router)

	port := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	s.Router.Logger.Fatal(s.Router.Start(port))
}
