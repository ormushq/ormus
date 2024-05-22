package main

import (
	"fmt"
	"log"

	"github.com/ormushq/ormus/adapter/scylladb/scyllainitialize"
	"github.com/ormushq/ormus/config"
)

func main() {
	cfg := config.C().Scylladb
	fmt.Println(cfg)
	// Create a new ScyllaDB connection instance
	conn := scyllainitialize.NewScyllaDBConnection(
		cfg.Consistency,
		cfg.Keyspace,
		cfg.Hosts...,
	)

	err := scyllainitialize.CreateKeySpace(
		cfg.Consistency,
		cfg.Keyspace,
		cfg.Hosts...,
	)
	if err != nil {
		log.Fatal("Failed to create ScyllaDB keyspace:", err)
	}

	// Get a ScyllaDB session using the created connection
	session, err := scyllainitialize.GetConnection(conn)
	if err != nil {
		log.Println("Failed to get ScyllaDB session:", err)
	}
	defer session.Close()

	err = scyllainitialize.RunMigrations(conn, "./source/repository/scylladb/")
	if err != nil {
		log.Println("Failed to run migrations:", err)
	}
}
