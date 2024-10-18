package sourceparam

type DeleteRequest struct {
	UserID   string `json:"-"`
	SourceID string `json:"-" param:"sourceID"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
