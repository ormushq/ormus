package projectstub

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/ormushq/ormus/manager/entity"
)

// TODO: this package have to be renamed to stub... because this logic is not mocking, but stubbing.

const RepoErr = "repository error"

type MockProject struct {
	projects []entity.Project
	hasErr   bool
}

func (m *MockProject) Create(project entity.Project) (entity.Project, error) {
	if m.hasErr {
		return entity.Project{}, fmt.Errorf(RepoErr)
	}

	project.ID = "new-id"
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	project.DeletedAt = nil
	project.Deleted = false

	m.projects = append(m.projects, project)

	return project, nil
}

func (m *MockProject) IsCreated(id string) bool {
	for _, val := range m.projects {
		if val.ID == id {
			return true
		}
	}

	return false
}

func (m *MockProject) GetWithID(id string) (entity.Project, error) {
	for _, val := range m.projects {
		if val.ID == id {
			return val, nil
		}
	}

	return entity.Project{}, gocql.ErrNotFound
}

func (m *MockProject) Update(project entity.Project) (entity.Project, error) {
	for id, val := range m.projects {
		if val.ID == project.ID {
			project.UpdatedAt = time.Now()

			m.projects[id] = project
		}
	}

	return project, nil
}

func (m *MockProject) Delete(project entity.Project) error {
	for id, val := range m.projects {
		if val.ID == project.ID {
			t := time.Now()
			project.DeletedAt = &t
			project.Deleted = true

			m.projects[id] = project
		}
	}

	return nil
}

func (m *MockProject) List(userID string, _ int64, limit int) ([]entity.Project, error) {
	li := []entity.Project{}
	count := 0
	for _, val := range m.projects {
		if val.Deleted && val.UserID == userID {
			count++
			li = append(li, val)
		}
		if count == limit {
			break
		}
	}

	return li, nil
}

func (m *MockProject) HaseMore(_ string, _ int64, _ int) (bool, error) {
	return false, nil
}

func New(hasErr bool) MockProject {
	const projectInMemoryDBSize = 0
	projects := make([]entity.Project, projectInMemoryDBSize)

	return MockProject{hasErr: hasErr, projects: projects}
}
