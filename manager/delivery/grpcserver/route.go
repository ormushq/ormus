package grpcserver

import (
	"fmt"
	"net"

	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/delivery/grpcserver/sourcehandler"
	"google.golang.org/grpc"
)

type SetupServices struct {
	WriteKeyValidationHandler sourcehandler.WriteKeyValidationHandler
}

type Server struct {
	WriteKeyValidationHandler sourcehandler.WriteKeyValidationHandler
	config                    manager.Config
}

func New(setupServices SetupServices, config manager.Config) *Server {
	return &Server{
		WriteKeyValidationHandler: setupServices.WriteKeyValidationHandler,
		config:                    config,
	}
}

func (s Server) Server() *grpc.Server {
	address := fmt.Sprintf(":%d", s.config.Application.GrpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	source.RegisterIsWriteKeyValidServer(grpcServer, s.WriteKeyValidationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

	return grpcServer
}
