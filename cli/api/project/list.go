package project

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) List(perPage, lastTokenID string) types.Request {
	return types.Request{
		Path:                  "projects",
		Method:                http.MethodGet,
		AuthorizationRequired: true,
		QueryParams: map[string]string{
			"last_token_id": lastTokenID,
			"per_page":      perPage,
		},
	}
}
