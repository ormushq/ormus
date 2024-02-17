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




