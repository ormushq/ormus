package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/mock"
	"github.com/ormushq/ormus/source/service/eventservice"
	"github.com/ormushq/ormus/source/validator/eventvalidator"
)

func main() {
	cfg := config.C()

	eventSvc, eventVld := setupServices(cfg)

	server := httpserver.New(cfg.Source, eventSvc, eventVld)
	server.Serve()
}

func setupServices(cfg config.Config) (eventSvc eventservice.Service, eventVld eventvalidator.Validator) {

	// TODO: add event Repo syclladb
	eventRepo := mock.NewMockRepository(false)

	eventSvc = eventservice.New(&eventRepo)
	eventVld = eventvalidator.New(eventRepo)

	return eventSvc, eventVld
}
