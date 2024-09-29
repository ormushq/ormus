package sourceparam

import "github.com/ormushq/ormus/manager/entity"

type CreateRequest struct {
	UserID      string `json:"-"`
	ProjectID   string `json:"project_id"`
	Name        string `json:"name"  example:"test"`
	Description string `json:"description"  example:"description"`
}

type CreateResponse struct {
	Source entity.Source `json:"source"`
}
