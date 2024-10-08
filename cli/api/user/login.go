package user

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Login(email, password string) types.Request {
	return types.Request{
		Path:                  "users/login",
		Method:                http.MethodPost,
		AuthorizationRequired: false,
		Body: map[string]string{
			"email":    email,
			"password": password,
		},
	}
}
