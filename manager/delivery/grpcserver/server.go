package grpcserver

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/contract/protobuf/manager/goproto/writekey"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/repository/sourcerepo"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type Server struct {
	writekey.UnimplementedWriteKeyManagerServer
	sourceRepo sourcerepo.SourceRepository
	config     manager.Config
	grpcServer *grpc.Server
	mu         sync.Mutex
}

func New(sourceRepo sourcerepo.SourceRepository, config manager.Config) Server {
	return Server{
		sourceRepo: sourceRepo,
		config:     config,
	}
}

func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	address := fmt.Sprintf(":%s", s.config.GRPCServiceAddress)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.grpcServer = grpc.NewServer()
	writekey.RegisterWriteKeyManagerServer(s.grpcServer, s)

	log.Printf("manager gRPC server listening at %s\n", address)
	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.grpcServer == nil {
		return nil
	}

	stopped := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		s.grpcServer.Stop()
	case <-stopped:
	}

	s.grpcServer = nil

	return nil
}
