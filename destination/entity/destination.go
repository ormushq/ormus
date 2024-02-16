package entity

// Destination represents where the processed event is sent.
type Destination struct {
	ID          string
	Type        string
	Name        string
	Description string
	Metadata    map[string]interface{}
}
