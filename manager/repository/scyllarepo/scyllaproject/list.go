package scyllaproject

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["List"] = scyllarepo.Statement{
		Query:  "SELECT id,token(id) as token_id, user_id, name, description, created_at, updated_at, deleted_at FROM projects where user_id = ?  and token(id) >  ?  and deleted = false  LIMIT ? ALLOW FILTERING;",
		Values: []string{"user_id", "last_token", "limit"},
	}
}

func (r Repository) List(userID string, lastToken int64, limit int) ([]entity.Project, error) {
	query, err := r.db.GetStatement(statements["List"])
	if err != nil {
		return nil, err
	}

	query.BindMap(qb.M{
		"user_id":    userID,
		"limit":      limit,
		"last_token": lastToken,
	})

	var projects []entity.Project
	if err = query.Select(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}
