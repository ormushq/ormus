package scyllauser

import (
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["Register"] = scyllarepo.Statement{
		Query:  "INSERT INTO users (id, created_at, updated_at, email, password, is_active) VALUES (?, ?, ?, ?, ?, ?)",
		Values: []string{"id", "created_at", "updated_at", "email", "password", "is_active"},
	}
}

func (r Repository) Register(u entity.User) (entity.User, error) {
	query, err := r.db.GetStatement(statements["Register"])
	if err != nil {
		return entity.User{}, err
	}

	u.ID = uuid.New().String()

	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt

	query.BindMap(qb.M{
		"id":         u.ID,
		"email":      u.Email,
		"password":   u.Password,
		"is_active":  u.IsActive,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	})

	err = query.Exec()
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}
