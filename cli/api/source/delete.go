package source

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Delete(sourceID string) types.Request {
	return types.Request{
		Path:                  "sources/%s",
		Method:                http.MethodDelete,
		AuthorizationRequired: true,
		URLParams: []any{
			sourceID,
		},
	}
}
