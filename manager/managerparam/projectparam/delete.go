package projectparam

type DeleteRequest struct {
	UserID    string `json:"-"`
	ProjectID string `json:"-" param:"projectID"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
