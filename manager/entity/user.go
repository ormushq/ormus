package entity

import "time"

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string
	Password  string
	IsActive  bool
	// TODO: do we need role?
}
