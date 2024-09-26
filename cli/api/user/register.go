package user

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Register(name, email, password string) types.Request {
	return types.Request{
		Path:                  "users/register",
		Method:                http.MethodPost,
		AuthorizationRequired: false,
		Body: map[string]string{
			"name":     name,
			"email":    email,
			"password": password,
		},
	}
}
