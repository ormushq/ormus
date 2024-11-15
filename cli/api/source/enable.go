package source

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Enable(sourceID string) types.Request {
	return types.Request{
		Path:                  "sources/%s/enable",
		Method:                http.MethodGet,
		AuthorizationRequired: true,
		Header:                nil,
		URLParams: []any{
			sourceID,
		},
	}
}
