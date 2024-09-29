package scyllaproject

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["GetWithId"] = scyllarepo.Statement{
		Query:  "SELECT id,token(id) as token_id, write_key, name, description, project_id, owner_id, status, created_at, updated_at, deleted_at FROM source where id = ? and deleted = false;",
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
