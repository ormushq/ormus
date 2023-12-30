package entity

type ConnectionType string

const (
	EventStream ConnectionType = "event-stream"
	Storage     ConnectionType = "storage"
	ReversETL   ConnectionType = "reverse-ETL"
)
