package param

// TODO: do we need project.description?

type CreateProjectRequest struct {
	UserID string
	Name   string
}

type CreateProjectResponse struct {
	ID   string
	Name string
}
