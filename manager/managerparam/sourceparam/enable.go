package sourceparam

type EnableRequest struct {
	UserID   string `json:"-"`
	SourceID string `json:"-" param:"SourceID"`
}

type EnableResponse struct {
	Message string `json:"message"`
}
