package scyllasource

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["GetWithId"] = scyllarepo.Statement{
		Query:  "SELECT id,token(id) as token_id, write_key, name, description, project_id, owner_id, status, created_at, updated_at, deleted_at FROM sources where id = ? and deleted = false  ALLOW FILTERING;",
		Values: []string{"id"},
	}
}

func (r Repository) GetWithID(id string) (entity.Source, error) {
	query, err := r.db.GetStatement(statements["GetWithId"])
	if err != nil {
		return entity.Source{}, err
	}

	query.BindMap(qb.M{
		"id": id,
	})

	var source entity.Source
	if err := query.Get(&source); err != nil {
		return entity.Source{}, err
	}

	return source, nil
}
