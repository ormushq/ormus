package scyllarepo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/scylladb/gocqlx/v2/qb"
)

func (a StorageAdapter) DoesUserExistsByEmail(email string) (bool, error) {
	var id string
	query := "SELECT id FROM users WHERE email = ? AND is_active = true LIMIT 1"
	names := []string{"email"}
	query1 := a.ScyllaConn.Query(query, names)
	query1.BindMap(qb.M{
		"email": email,
	})

	found := query1.Iter().Scan(&id)
	if err := query1.Iter().Close(); err != nil {
		logger.L().Debug("Error closing iterator", "err msg:", err)

		return false, err
	}

	logger.L().Debug("Query executed successfully", "is found:", found)

	return found, nil
}

func (a StorageAdapter) Register(u entity.User) (*entity.User, error) {
	u.ID = uuid.New().String()
	var deletedAtValue string
	if u.DeletedAt == nil {
		deletedAtValue = "NULL"
	} else {
		deletedAtValue = fmt.Sprintf("'%s'", u.DeletedAt.Format(time.RFC3339))
	}

	query := fmt.Sprintf("INSERT INTO users (id, created_at, updated_at, deleted_at, email, password, is_active) VALUES ('%s', '%s', '%s', %s, '%s', '%s', %t)",
		u.ID,
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		deletedAtValue,
		u.Email,
		u.Password,
		u.IsActive,
	)

	err := a.ScyllaConn.ExecStmt(query)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (a StorageAdapter) GetUserByEmail(email string) (*entity.User, error) {
	// TODO : check if this is the right way to do it
	stmt := "SELECT * FROM users WHERE email = ? AND is_active = true"

	names := []string{"email"}

	query := a.ScyllaConn.Query(stmt, names)

	query.BindMap(qb.M{
		"email": email,
	})

	var user entity.User
	if err := query.Get(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
