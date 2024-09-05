package param

type RegisterRequest struct {
	Name     string `json:"name" example:"name"`
	Email    string `json:"email" example:"name@test.com"`
	Password string `json:"password" example:"123Qwe!@#"`
}

type RegisterResponse struct {
	Email string `json:"email" example:"name@test.com"`
	ID    string `json:"id" example:"f90631e0-aad3-4eb1-8cef-1478711e16e9"`
}
