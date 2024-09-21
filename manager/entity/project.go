package entity

import (
	"time"
)

// Project is the main object for managing different connections.
type Project struct {
	ID          string     `json:"id"`
	TokenID     string     `json:"token_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Deleted     bool       `json:"deleted"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	UserID      string     `json:"user_id"`
}
