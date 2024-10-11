package scyllasource

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Delete"] = scyllarepo.Statement{
		Query:  "update sources set deleted_at = ?,deleted = ?,status = ? where id = ?",
		Values: []string{"deleted_at", "deleted", "status", "id"},
	}
}

func (r Repository) Delete(source entity.Source) error {
	query, err := r.db.GetStatement(statements["Delete"])
	if err != nil {
		return err
	}
	t := time.Now()
	source.DeletedAt = &t

	query.BindMap(qb.M{
		"id":         source.ID,
		"deleted":    true,
		"status":     entity.SourceStatusNotActive,
		"deleted_at": source.DeletedAt,
	})

	err = query.Exec()

	return err
}
