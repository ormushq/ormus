package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/delivery/httpserver/statushandler"
)

func main() {
	handlers := []httpserver.Handler{
		statushandler.New(),
	}

	httpServer := httpserver.New(config.C().Source, handlers)

	httpServer.Serve()
}
