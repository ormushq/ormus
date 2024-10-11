package scyllasource

import (
	"log/slog"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["List"] = scyllarepo.Statement{
		Query:  "SELECT id,token(id) as token_id, write_key, name, description, project_id, owner_id, status, created_at, updated_at, deleted_at FROM sources where owner_id = ? AND token(id) >  ?  AND deleted = false  LIMIT ? ALLOW FILTERING;",
		Values: []string{"owner_id", "last_token", "limit"},
	}
}

func (r Repository) List(ownerID string, lastToken int64, limit int) ([]entity.Source, error) {
	query, err := r.db.GetStatement(statements["List"])
	if err != nil {
		return nil, err
	}

	query.BindMap(qb.M{
		"owner_id":   ownerID,
		"limit":      limit,
		"last_token": lastToken,
	})

	var sources []entity.Source
	if err := query.Select(&sources); err != nil {
		logger.L().Error(err.Error(), slog.String("query", query.String()))

		return nil, err
	}

	return sources, nil
}
