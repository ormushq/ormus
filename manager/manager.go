package manager

type Repository interface{}

// Manager is a service to srcconfig pipelines/connections.
type Manager struct {
	repo Repository
}

func New(repo Repository) Manager {
	return Manager{repo: repo}
}
