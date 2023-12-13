package param

import "time"

type UserInfo struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Token    `json:"token"`
}
