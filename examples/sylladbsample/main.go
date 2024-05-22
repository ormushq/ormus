package main

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/adapter/scylladb"
)

func main() {
	// Set up ScyllaDB configuration
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
		log.Fatal("Error creating ScyllaDB session:", err)
	}
	defer session.Close()

	// Use the session to interact with ScyllaDB

}
