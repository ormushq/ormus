package projectparam

import "github.com/ormushq/ormus/manager/entity"

type ListRequest struct {
	UserID      string
	LastTokenID string `query:"last_token_id"`
	PerPage     int    `query:"per_page"`
}

type ListResponse struct {
	Projects    []entity.Project `json:"projects"`
	LastTokenID string           `json:"last_token"`
	PerPage     int              `json:"per_page"`
	HasMore     bool             `json:"has_more"`
}
