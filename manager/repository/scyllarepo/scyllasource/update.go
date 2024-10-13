package scyllasource

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Update"] = scyllarepo.Statement{
		Query:  "update sources set name = ?,description = ?, status = ?, updated_at = ? where id = ?",
		Values: []string{"name", "description", "status", "updated_at", "id"},
	}
}

func (r Repository) Update(source entity.Source) (entity.Source, error) {
	query, err := r.db.GetStatement(statements["Update"])
	if err != nil {
		return entity.Source{}, err
	}
	source.UpdatedAt = time.Now()

	query.BindMap(qb.M{
		"id":          source.ID,
		"name":        source.Name,
		"description": source.Description,
		"status":      source.Status,
		"updated_at":  source.UpdatedAt,
	})

	if err := query.Exec(); err != nil {
		return entity.Source{}, err
	}

	return source, nil
}
