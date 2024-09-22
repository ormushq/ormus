package scyllaproject

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Update"] = scyllarepo.Statement{
		Query:  "update projects set name = ?,description = ?, updated_at = ? where id = ?",
		Values: []string{"name", "description", "updated_at", "id"},
	}
}

func (r Repository) Update(project entity.Project) (entity.Project, error) {
	query, err := r.db.GetStatement(statements["Update"])
	if err != nil {
		return entity.Project{}, err
	}
	project.UpdatedAt = time.Now()

	query.BindMap(qb.M{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
		"updated_at":  project.UpdatedAt,
	})

	if err := query.Exec(); err != nil {
		return entity.Project{}, err
	}

	return project, nil
}
