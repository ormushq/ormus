package scyllaproject

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Delete"] = scyllarepo.Statement{
		Query:  "update projects set deleted_at = ?,deleted = ? where id = ?",
		Values: []string{"deleted_at", "deleted", "id"},
	}
}

func (r Repository) Delete(project entity.Project) error {
	query, err := r.db.GetStatement(statements["Delete"])
	if err != nil {
		return err
	}
	t := time.Now()
	project.DeletedAt = &t

	query.BindMap(qb.M{
		"id":         project.ID,
		"deleted":    true,
		"deleted_at": project.DeletedAt,
	})

	err = query.Exec()

	return err
}
