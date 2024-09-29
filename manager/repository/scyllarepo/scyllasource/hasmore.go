package scyllaproject

import (
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["HaseMore"] = scyllarepo.Statement{
		Query:  "SELECT COUNT(id) as total FROM source where owner_id = ? and token(id) > ?;",
		Values: []string{"owner_id", "last_token"},
	}
}

func (r Repository) HaseMore(ownerID, lastToken string, perPage int) (bool, error) {
	query, err := r.db.GetStatement(statements["HaseMore"])
	if err != nil {
		return false, err
	}

	query.BindMap(qb.M{
		"owner_id":   ownerID,
		"last_token": lastToken,
	})

	var total int
	if err := query.Scan(&total); err != nil {
		return false, err
	}

	return (total - perPage) > 0, nil
}
