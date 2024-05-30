package httpserver

import (
	"fmt"
	"github.com/ormushq/ormus/source/delivery/httpserver/eventhandler"
	"github.com/ormushq/ormus/source/delivery/httpserver/statushandler"
	"github.com/ormushq/ormus/source/delivery/httpserver/userhandler"
	"github.com/ormushq/ormus/source/service/eventservice"
	"github.com/ormushq/ormus/source/validator/eventvalidator"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source"
)

// Server is the main object for managing http configurations and handlers.
type Server struct {
	config        source.Config
	Router        *echo.Echo
	userhandler   userhandler.Handler
	eventHandler  eventhandler.Handler
	statusHandler statushandler.Handler
}

// New Set up a new server object.
// Setup a new server object.
func New(c source.Config, eventSvc eventservice.Service, eventVld eventvalidator.Validator) Server {
	return Server{
		config: source.Config{
			HTTPServer: c.HTTPServer,
		},
		Router:        echo.New(),
		eventHandler:  eventhandler.New(eventSvc, eventVld),
		statusHandler: statushandler.New(),
	}
}

// Serve Start server connection.
func (s Server) Serve() {
	s.userhandler.SetRoutes(s.Router)
	s.eventHandler.SetEventRoute(s.Router)
	s.statusHandler.SetRoutes(s.Router)

	port := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	s.Router.Logger.Fatal(s.Router.Start(port))
}
