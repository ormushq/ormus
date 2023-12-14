package error

import "fmt"

var (
	ErrWrongCredentials = fmt.Errorf("username or password isn't correct")
)
