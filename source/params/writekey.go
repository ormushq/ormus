package params

type WriteKey struct {
	OwnerID   string `json:"owner_id"`
	ProjectID string `json:"project_id"`
	WriteKey  string `json:"write_key"`
}
