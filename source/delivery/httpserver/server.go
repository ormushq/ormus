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

func New() Server {
	return Server{
		Router: echo.New(),
	}
}

func (s Server) Serve() {
	s.userhandler.SetRoutes(s.Router)

	c := config.Load()

	port := fmt.Sprintf(":%d", c.HTTPServer.Port)
	s.Router.Logger.Fatal(s.Router.Start(port))
}
