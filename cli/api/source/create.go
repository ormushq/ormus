package source

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Create(name, description, projectID string) types.Request {
	return types.Request{
		Path:                  "sources",
		Method:                http.MethodPost,
		AuthorizationRequired: true,
		Header:                nil,
		Body: map[string]string{
			"name":        name,
			"description": description,
			"project_id":  projectID,
		},
	}
}
