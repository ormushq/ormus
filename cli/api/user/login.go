package user

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Login(email, password string) types.Request {
	return types.Request{
		Path:                  "login",
		Method:                http.MethodPost,
		AuthorizationRequired: false,
		Header:                nil,
		Body: map[string]string{
			email:    email,
			password: password,
		},
	}
}
