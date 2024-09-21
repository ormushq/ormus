package projectparam

import "github.com/ormushq/ormus/manager/entity"

type UpdateRequest struct {
	UserID      string `json:"-"`
	ProjectID   string `json:"-" param:"projectID"`
	Name        string `json:"name" example:"name"`
	Description string `json:"description"  example:"description"`
}

type UpdateResponse struct {
	Project entity.Project `json:"project"`
}
