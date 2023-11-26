package entity

type Pipeline struct {
	// TODO: should we add integrations inside Pipeline?
	Integrations []Integration
}

// Connection is an object for configure pipeline.
type Connection struct {
	Pipes        []Pipeline
	Integrations []Integration
}

// TODO: I'm not sure integration is right name for destination object.
type Integration struct{}

// TODO: Do we need source object?
type Source struct{}
