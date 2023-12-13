package error

import "fmt"

var (
	JwtEmptyUserErr = fmt.Errorf("for generating a JWT token, email is required")
)
