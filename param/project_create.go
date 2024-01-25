package param

// TODO: do we need project.description?

type CreateProjectRequest struct {
	Name   string
	UserID string
}

type CreateProjectResponse struct {
	ID   string
	Name string
}
