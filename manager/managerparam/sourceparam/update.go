package sourceparam

import "github.com/ormushq/ormus/manager/entity"

type UpdateRequest struct {
	UserID      string `json:"-"`
	SourceID    string `json:"-" param:"SourceID"`
	Name        string `json:"name" example:"name"`
	Description string `json:"description"  example:"description"`
	Status      string `json:"status"  example:"active"`
}

type UpdateResponse struct {
	Project entity.Project `json:"project"`
}
