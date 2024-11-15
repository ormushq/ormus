package sourceparam

type DisableRequest struct {
	UserID   string `json:"-"`
	SourceID string `json:"-" param:"SourceID"`
}

type DisableResponse struct {
	Message string `json:"message"`
}
