package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/delivery/httpserver"
)

func main() {
	config := config.New(config.Option{
		YamlFilePath: "config.yml",
	})

	server := httpserver.New(config.Source)

	server.Serve()
}
