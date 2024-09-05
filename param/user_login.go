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
	Email    string `json:"email"  example:"name@test.com"`
	Password string `json:"password"  example:"123Qwe!@#"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Token    `json:"token"`
}
