package scyllaproject

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["IsProjectExist"] = scyllarepo.Statement{
		Query:  "SELECT id FROM projects WHERE id = ? LIMIT 1",
		Values: []string{"id"},
	}
}

func (r Repository) Exist(projectID string) (bool, error) {
	var id string
	query, err := r.db.GetStatement(statements["IsProjectExist"])
	if err != nil {
		return false, err
	}
	query.BindMap(qb.M{
		"id": projectID,
	})

	found := query.Iter().Scan(&id)
	if err = query.Iter().Close(); err != nil {
		logger.L().Debug("Error closing iterator", "err msg:", err)

		return false, err
	}

	logger.L().Debug("Query executed successfully", "is found:", found)

	return found, nil
}
