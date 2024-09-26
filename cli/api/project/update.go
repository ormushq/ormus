package project

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Update(projectID, name, description string) types.Request {
	return types.Request{
		Path:                  "projects/%s",
		Method:                http.MethodPost,
		AuthorizationRequired: true,
		Header:                nil,
		URLParams: []any{
			projectID,
		},
		Body: map[string]string{
			"name":        name,
			"description": description,
		},
	}
}
