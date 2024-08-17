package user

import (
	"github.com/ormushq/ormus/cli/api/types"
	"net/http"
)

func (c Client) Login(email string, password string) types.Request {
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
