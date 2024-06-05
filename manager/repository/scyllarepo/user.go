package scyllarepo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/scylladb/gocqlx/v2/qb"
)

func (a StorageAdapter) DoesUserExistsByEmail(email string) (bool, error) {
	// if i dont use ALLOW FILTERING, it will return an error:
	// Cannot execute this query as it might involve data filtering and thus may have unpredictable performance.
	// If you want to execute this query despite the performance unpredictability, use ALLOW FILTERING
	// what should do is to use a query like this:
	// TODO : check if this is the right way to do it
	// ALLOW FILTERING is a mechanism provided by ScyllaDB and Apache Cassandra to allow querying on non-indexed columns at the expense of performance.
	// It's typically used as a last resort and should be avoided whenever possible,
	// especially in production environments with large datasets,
	// as it can lead to degraded performance and increased load on the database.

	stmt := "SELECT COUNT(*) FROM users WHERE email = ? AND is_active = true ALLOW FILTERING"
	// Initialize the names of placeholders in your query
	names := []string{"email"}

	// Execute the query
	query := a.ScyllaConn.Query(stmt, names)
	// Bind the parameters
	query.BindMap(qb.M{
		"email": email,
	})
	// Execute the query
	var count int
	if err := query.Scan(&count); err != nil {
		return false, err
	}

	// Check if count is greater than 0, indicating user exists
	exists := count > 0

	return exists, nil
}

func (a StorageAdapter) Register(u entity.User) (*entity.User, error) {
	u.ID = uuid.New().String()
	var deletedAtValue string
	if u.DeletedAt == nil {
		deletedAtValue = "NULL"
	} else {
		deletedAtValue = fmt.Sprintf("'%s'", u.DeletedAt.Format(time.RFC3339))
	}
	// Now it's true by default until our authentication system works properly.
	// TODO: The user is set to active when he has confirmed his email
	u.IsActive = true
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
	stmt := "SELECT id, created_at, updated_at, deleted_at, email, password, is_active FROM users WHERE email = ? AND is_active = true ALLOW FILTERING"

	// Initialize the names of placeholders in your query
	names := []string{"email"}

	// Execute the query
	query := a.ScyllaConn.Query(stmt, names)
	// Bind the parameters
	query.BindMap(qb.M{
		"email": email,
	})
	// Execute the query

	// Execute the query
	var user entity.User
	if err := query.Get(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
