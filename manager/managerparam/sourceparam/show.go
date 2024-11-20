package sourceparam

import "github.com/ormushq/ormus/manager/entity"

type ShowRequest struct {
	UserID   string `json:"-"`
	SourceID string `json:"-" param:"SourceID"`
}

type ShowResponse struct {
	Source entity.Source `json:"source"`
}
