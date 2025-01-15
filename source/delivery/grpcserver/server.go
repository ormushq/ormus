package grpcserver

import (
	"fmt"
	"net"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/source"
	"google.golang.org/grpc"
)

// Server is the main object for managing grpc configurations and handlers.
type Server struct {
	Config source.Config
}

func New() Server {
	return Server{}
}

// Start server connection.
func (s Server) Start() {
	// config the grpc port and network type
	address := fmt.Sprintf(":%d", s.Config.HTTPServer.Port)
	listener, err := net.Listen(s.Config.HTTPServer.Network, address)
	if err != nil {
		logger.L().Error(err.Error())
	}

	grpcserver := grpc.NewServer()

	// start the grpc server
	if err := grpcserver.Serve(listener); err != nil {
		logger.L().Error(err.Error())
	}
}
