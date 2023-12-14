package error

import "fmt"

var (
	ErrJwtEmptyUser = fmt.Errorf("for generating a JWT token, email is required")
)
