package sourceparam

import "github.com/ormushq/ormus/manager/entity"

type UpdateRequest struct {
	UserID      string `json:"-"`
	SourceID    string `json:"-" param:"SourceID"`
	Name        string `json:"name" example:"updated name"`
	Description string `json:"description"  example:"updated description"`
	Status      string `json:"status"  example:"active"`
}

type UpdateResponse struct {
	Source entity.Source `json:"source"`
}
