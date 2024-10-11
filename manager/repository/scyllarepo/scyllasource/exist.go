package scyllasource

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["IsSourceExist"] = scyllarepo.Statement{
		Query:  "SELECT id FROM sources WHERE id = ? LIMIT 1",
		Values: []string{"id"},
	}
}

func (r Repository) IsExist(sourceID string) (bool, error) {
	var id string
	query, err := r.db.GetStatement(statements["IsSourceExist"])
	if err != nil {
		return false, err
	}
	query.BindMap(qb.M{
		"id": sourceID,
	})

	found := query.Iter().Scan(&id)
	if err = query.Iter().Close(); err != nil {
		logger.L().Debug("Error closing iterator", "err msg:", err)

		return false, err
	}

	logger.L().Debug("Query executed successfully", "is found:", found)

	return found, nil
}
