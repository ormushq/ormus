package entity

import "github.com/ormushq/ormus/destination/entity"

type Pipeline struct {
	// TODO: should we add integrations inside Pipeline?
	Integrations []entity.Integration
}

// Connection is an object for configure pipeline.
type Connection struct {
	Pipes        []Pipeline
	Integrations []entity.Integration
}

// TODO: Do we need source object?
type Source struct{}
