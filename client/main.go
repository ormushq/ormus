package main

import "github.com/ormushq/ormus/source/delivery/httpserver"

func main() {
	server := httpserver.New()

	server.Serve()
}