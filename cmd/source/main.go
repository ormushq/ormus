package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/delivery/httpserver"
	"github.com/ormushq/ormus/source/repository/scylladb/event"
	"github.com/ormushq/ormus/source/service/eventservice"
	"github.com/ormushq/ormus/source/validator/eventvalidator"
	"log/slog"
	"time"
)

func main() {
	cfg := config.C()

	eventSvc, eventVld := setupServices(cfg)

	server := httpserver.New(cfg.Source, eventSvc, eventVld)
	server.Serve()
}

func setupServices(cfg config.Config) (eventSvc eventservice.Service, eventVld eventvalidator.Validator) {

	// TODO: add event Repo syclladb
	//eventRepo := mock.NewMockRepository(false)

	config := scylladb.Config{
		Hosts:          []string{"127.0.0.1:9042"},
		Consistency:    gocql.One,
		Keyspace:       "your_keyspace",
		TimeoutCluster: 5 * time.Second,
		NumRetries:     5,
		MinRetryDelay:  time.Second,
		MaxRetryDelay:  10 * time.Second,
	}

	// Create a new ScyllaDB session
	session, err := scylladb.New(config)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error creating ScyllaDB session", err))
	}
	sycllaRepo := event.New(session)
	eventSvc = eventservice.New(&sycllaRepo)
	eventVld = eventvalidator.New(sycllaRepo)

	return eventSvc, eventVld
}
