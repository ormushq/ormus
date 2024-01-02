package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source"
	"github.com/ormushq/ormus/source/delivery/httpserver/userhandler"
)

// Server is the main object for managing http configurations and handlers.
type Server struct {
	config      source.Config
	Router      *echo.Echo
	userhandler userhandler.Handler
}

// Setup a new server object.
// New Set up a new server object.
func New(c source.Config) Server {
	return Server{
		config: source.Config{
			HTTPServer: c.HTTPServer,
		},
		Router: echo.New(),
	}
}

// Start server connection.
func (s Server) Serve() {
	s.userhandler.SetRoutes(s.Router)

	port := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	s.Router.Logger.Fatal(s.Router.Start(port))
}
