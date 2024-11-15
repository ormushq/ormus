package source

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Update(sourceID, name, description string) types.Request {
	return types.Request{
		Path:                  "sources/%s",
		Method:                http.MethodPost,
		AuthorizationRequired: true,
		Header:                nil,
		URLParams: []any{
			sourceID,
		},
		Body: map[string]string{
			"name":        name,
			"description": description,
		},
	}
}
