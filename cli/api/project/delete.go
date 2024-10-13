package project

import (
	"net/http"

	"github.com/ormushq/ormus/cli/api/types"
)

func (c Client) Delete(projectID string) types.Request {
	return types.Request{
		Path:                  "projects/%s",
		Method:                http.MethodDelete,
		AuthorizationRequired: true,
		URLParams: []any{
			projectID,
		},
	}
}
