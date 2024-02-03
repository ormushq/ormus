package projectmock

import (
	"fmt"
	"github.com/ormushq/ormus/manager/entity"
	"time"
)

// TODO: this package have to be renamed to stub... because this logic is not mocking, but stubbing.

const RepoErr = "repository error"

type MockProject struct {
	projects []entity.Project
	hasErr   bool
}

func (m MockProject) Create(name string, email string) (entity.Project, error) {
	if m.hasErr {
		return entity.Project{}, fmt.Errorf(RepoErr)
	}

	project := entity.Project{
		ID:          "new-id",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
		Name:        name,
		Description: "",
		UserID:      fmt.Sprintf("a user with this email: %s", email),
	}

	m.projects = append(m.projects, project)
	return project, nil
}

func New(hasErr bool) MockProject {
	projects := make([]entity.Project, 10)
	return MockProject{hasErr: hasErr, projects: projects}
}
