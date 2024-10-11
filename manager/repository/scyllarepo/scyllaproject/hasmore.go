package scyllaproject

import (
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["HaseMore"] = scyllarepo.Statement{
		Query:  "SELECT COUNT(id) as total FROM projects where user_id = ? and token(id) > ?  ALLOW FILTERING;",
		Values: []string{"user_id", "last_token"},
	}
}

func (r Repository) HaseMore(userID string, lastToken int64, perPage int) (bool, error) {
	query, err := r.db.GetStatement(statements["HaseMore"])
	if err != nil {
		return false, err
	}

	query.BindMap(qb.M{
		"user_id":    userID,
		"last_token": lastToken,
	})

	var total int
	if err := query.Scan(&total); err != nil {
		return false, err
	}

	return (total - perPage) > 0, nil
}
