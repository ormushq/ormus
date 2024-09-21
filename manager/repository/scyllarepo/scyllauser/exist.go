package scyllauser

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["DoesUserExistsByEmail"] = scyllarepo.Statement{
		Query:  "SELECT id FROM users WHERE email = ? AND is_active = true LIMIT 1",
		Values: []string{"email"},
	}
}

func (r Repository) DoesUserExistsByEmail(email string) (bool, error) {
	var id string
	query, err := r.db.GetStatement(statements["DoesUserExistsByEmail"])
	if err != nil {
		return false, err
	}
	query.BindMap(qb.M{
		"email": email,
	})

	found := query.Iter().Scan(&id)
	if err = query.Iter().Close(); err != nil {
		logger.L().Debug("Error closing iterator", "err msg:", err)

		return false, err
	}

	logger.L().Debug("Query executed successfully", "is found:", found)

	return found, nil
}
