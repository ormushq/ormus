package manager

type Repository interface{}

// Manager is a authservice to config pipelines/connections.
type Manager struct {
	repo Repository
}

func New(repo Repository) Manager {
	return Manager{repo: repo}
}
