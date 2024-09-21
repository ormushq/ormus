package projectparam

import "github.com/ormushq/ormus/manager/entity"

type CreateRequest struct {
	UserID      string `json:"-"`
	Name        string `json:"name" example:"name"`
	Description string `json:"description"  example:"description"`
}
type CreateThoughChannel struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateResponse struct {
	Project entity.Project `json:"project"`
}
