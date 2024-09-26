package project

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Create(name, description string) types.Request {
	return types.Request{
		Path:                  "projects",
		Method:                http.MethodPost,
		AuthorizationRequired: true,
		Header:                nil,
		Body: map[string]string{
			"name":        name,
			"description": description,
		},
	}
}
