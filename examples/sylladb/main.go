package main

import (
	"fmt"
	"log"

	"github.com/ormushq/ormus/adapter/scylladb"
	"github.com/ormushq/ormus/config"
)

func main() {
	fmt.Println("Hi")
	cfg := config.C()
	fmt.Println(cfg.Scylladb)
	// Create a new ScyllaDB session
	session, err := scylladb.New(cfg.Scylladb)
	if err != nil {
		log.Fatal("Error creating ScyllaDB session:", err)
	}
	defer session.Close()

	// Use the session to interact with ScyllaDB

}
