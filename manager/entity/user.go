package entity

import "time"

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	IsActive  bool       `json:"is_active"`
	// TODO: do we need role?
}
