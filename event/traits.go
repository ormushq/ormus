package event

import (
	"net/url"
	"time"
)

type Traits struct {
	CustomTraits CustomData
}

type UserTraits struct {
	Traits

	ID       string // Unique ID in your database for a user
	Industry string // Industry a user works in

	FirstName   string
	LastName    string
	Name        string // Full name of a user. If you only pass a first and last name Segment automatically fills in the full name for you.
	PhoneNumber string
	Username    string

	Title string // Title of a user, usually related to their position at a specific company.

	Website url.URL // User's website
	Address Address

	Age      int
	Avatar   url.URL // URL to an avatar image for the user.\
	Birthday time.Time

	Company Company

	Plan string //  Plan that a user is in. example(enterprise)

	CreatedAt time.Time

	Description string // Description of user, such as bio.
	Email       string
	Gender      string
}

type Company struct {
	Name          string
	ID            string // Unique ID in your database for a company
	EmployeeCount int
}

type Address struct {
	City       string
	Country    string
	PostalCode string
	State      string
	Street     string
}
