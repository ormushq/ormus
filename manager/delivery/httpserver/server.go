package httpserver

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/delivery/httpserver/projecthandler"
	"github.com/ormushq/ormus/manager/delivery/httpserver/sourcehandler"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
)

type SetupServicesResponse struct {
	UserHandler    userhandler.Handler
	ProjectHandler projecthandler.Handler
	SourceHandler  sourcehandler.Handler
}

type Server struct {
	config         manager.Config
	userHandler    userhandler.Handler
	sourceHandler  sourcehandler.Handler
	projectHandler projecthandler.Handler
	Router         *echo.Echo
}

func New(cfg manager.Config, setupSvc SetupServicesResponse) *Server {
	return &Server{
		config:         cfg,
		userHandler:    setupSvc.UserHandler,
		projectHandler: setupSvc.ProjectHandler,
		sourceHandler:  setupSvc.SourceHandler,
		Router:         echo.New(),
	}
}

func (s *Server) Server() {
	e := s.Router

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// TODO add this to config
		AllowOrigins: []string{"http://*.ormus.local"},
	}))

	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogURIPath:       true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			logger.L().
				Info("http-server",
					slog.String("RequestID", v.RequestID),
					slog.String("Host", v.Host),
					slog.String("ContentLength", v.ContentLength),
					slog.String("Protocol", v.Protocol),
					slog.String("Method", v.Method),
					slog.String("Latency", v.Latency.String()),
					slog.String("errMsg", errMsg),
					slog.String("RemoteIP", v.RemoteIP),
					slog.Int64("ResponseSize", v.ResponseSize),
					slog.String("URI", v.URI),
					slog.String("URIPath", v.URIPath),
					slog.Int("Status", v.Status),
				)

			return nil
		},
	}))

	s.userHandler.SetUserRoute(e)
	s.sourceHandler.SetRoutes(e)
	s.projectHandler.SetRoutes(e)

	e.GET("/health-check", s.healthCheck)

	routes, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Printf(err.Error())
	}
	fmt.Println(string(routes))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.Application.Port)))
}
