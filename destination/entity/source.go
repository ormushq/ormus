package entity

// Source is a website, server library, mobile SDK, or cloud application which can send data into Segment.
type Source struct {
	ID          string
	WriteKey    string
	Type        string
	Name        string
	Description string
	Metadata    map[string]interface{}
}
