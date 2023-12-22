package entity

import "time"

type Pipeline struct {
	// TODO: should we add integrations inside Pipeline?
	Integrations []Integration
}

// Connection is an object for configure pipeline.
type Connection struct {
	Pipes        []Pipeline
	Integrations []Integration
}

type Category struct {
	ID   int
	Name string
}

type ConnectionType string

const (
	EventStream ConnectionType = "event-stream"
	Storage     ConnectionType = "storage"
	ReversETL   ConnectionType = "reverse-ETL"
)

type Integration struct {
	Name             string
	CategoryID       Category
	Status           bool
	Source           Source
	Type             string
	ConnectionType   ConnectionType
	CreatedAt        *time.Time
	LatestSyncStatus *time.Time

	// TODO: Do we have write key field here?
	// TODO: What else do we need? is "configurations map[string]interface" good choice for other requirements?
}

// TODO: Do we need source object?
type Source struct{}
