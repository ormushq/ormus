package scyllaproject

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["GetWithId"] = scyllarepo.Statement{
		Query:  "SELECT id,token(id) as token_id, user_id, name, description, created_at, updated_at, deleted_at FROM projects where id = ?;",
		Values: []string{"id"},
	}
}

func (r Repository) GetWithID(id string) (entity.Project, error) {
	query, err := r.db.GetStatement(statements["GetWithId"])
	if err != nil {
		return entity.Project{}, err
	}

	query.BindMap(qb.M{
		"id": id,
	})

	var project entity.Project
	if err := query.Get(&project); err != nil {
		return entity.Project{}, err
	}

	return project, nil
}
