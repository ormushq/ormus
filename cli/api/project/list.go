package project

import (
	"github.com/ormushq/ormus/cli/api/types"
	"net/http"
)

func (c Client) List(perPage, lastTokenId string) types.Request {
	return types.Request{
		Path:                  "projects",
		Method:                http.MethodGet,
		AuthorizationRequired: true,
		Header:                nil,
		QueryParams: map[string]string{
			"lat_token_id": lastTokenId,
			"per_page":     perPage,
		},
		Body: map[string]string{},
	}
}
