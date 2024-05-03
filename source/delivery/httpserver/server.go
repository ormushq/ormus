package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source"
)

type Handler interface {
	SetRoutes(router *echo.Echo)
}

// Server is the main object for managing http configurations and handlers.
type Server struct {
	config   source.Config
	Router   *echo.Echo
	handlers []Handler
}

// Setup a new server object.
// New Set up a new server object.
// Setup a new server object.
func New(c source.Config, handlers []Handler) Server {
	return Server{
		config: source.Config{
			HTTPServer: c.HTTPServer,
		},
		Router:   echo.New(),
		handlers: handlers,
	}
}

// Run server connection.
func (s Server) Serve() {
	for _, h := range s.handlers {
		h.SetRoutes(s.Router)
	}

	port := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	s.Router.Logger.Fatal(s.Router.Start(port))
}
