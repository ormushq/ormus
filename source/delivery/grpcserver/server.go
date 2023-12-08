package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/ormushq/ormus/source/config"
	"google.golang.org/grpc"
)

type Server struct {
	Config config.Config
}

func New() Server {
	return Server{}
}

func (s Server) Start() {
	// config the grpc port and network type
	address := fmt.Sprintf(":%d", s.Config.HTTPServer.Port)
	listener, err := net.Listen(s.Config.HTTPServer.Network, address); if err != nil {
		log.Fatal(err)
	}

	grpcserver := grpc.NewServer()

	// start the grpc server
	if err := grpcserver.Serve(listener); err != nil {
		log.Fatal(err)
	}
}