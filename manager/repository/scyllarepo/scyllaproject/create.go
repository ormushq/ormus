package scyllaproject

import (
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Create"] = scyllarepo.Statement{
		Query:  "Insert into projects(id,name,description,user_id,created_at,updated_at, deleted)values(?,?,?,?,?,?,?)",
		Values: []string{"id", "name", "description", "user_id", "created_at", "updated_at", "deleted"},
	}
}

func (r Repository) Create(project entity.Project) (entity.Project, error) {
	query, err := r.db.GetStatement(statements["Create"])
	if err != nil {
		return entity.Project{}, err
	}
	project.ID = uuid.New().String()
	project.CreatedAt = time.Now()
	project.UpdatedAt = project.CreatedAt
	project.DeletedAt = nil
	query.BindMap(qb.M{
		"id":          project.ID,
		"user_id":     project.UserID,
		"name":        project.Name,
		"deleted":     false,
		"description": project.Description,
		"created_at":  project.CreatedAt,
		"updated_at":  project.UpdatedAt,
	})

	if err := query.Exec(); err != nil {
		return entity.Project{}, err
	}

	return project, nil
}
