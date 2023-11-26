package entity

import "time"

// Project is the main object for managing different connections.
type Project struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Description string
	UserID      string
}
