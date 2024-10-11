package scyllasource

import (
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Create"] = scyllarepo.Statement{
		Query:  "Insert into sources(id, write_key, name, description, project_id, owner_id, status, created_at,updated_at, deleted)values(?,?,?,?,?,?,?,?,?,?)",
		Values: []string{"id", "write_key", "name", "description", "project_id", "owner_id", "status", "created_at", "updated_at", "deleted"},
	}
}

func (r Repository) Create(source entity.Source) (entity.Source, error) {
	query, err := r.db.GetStatement(statements["Create"])
	if err != nil {
		return entity.Source{}, err
	}
	source.ID = uuid.New().String()
	source.CreatedAt = time.Now()
	source.UpdatedAt = source.CreatedAt
	source.DeletedAt = nil
	query.BindMap(qb.M{
		"id":          source.ID,
		"write_key":   source.WriteKey,
		"name":        source.Name,
		"description": source.Description,
		"project_id":  source.ProjectID,
		"owner_id":    source.OwnerID,
		"status":      entity.SourceStatusNotActive,
		"created_at":  source.CreatedAt,
		"updated_at":  source.UpdatedAt,
		"deleted":     false,
	})

	if err := query.Exec(); err != nil {
		return entity.Source{}, err
	}

	return source, nil
}
