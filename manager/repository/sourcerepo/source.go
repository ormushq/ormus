package sourcerepo

import (
	"context"
	"errors"
	"strings"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

var ErrWriteKeyNotFound = errors.New("write key not found")

type SourceRepository interface {
	GetWriteKey(ctx context.Context, key string) (entity.WriteKeyMetaData, error)
}

type sourceRepo struct {
	scyllaAdapter *scyllarepo.StorageAdapter
}

func New(scyllaAdapter *scyllarepo.StorageAdapter) SourceRepository {
	return &sourceRepo{scyllaAdapter: scyllaAdapter}
}

func (s sourceRepo) GetWriteKey(ctx context.Context,
	writeKey string,
) (entity.WriteKeyMetaData, error) {
	var metadata entity.WriteKeyMetaData

	stmt, names := qb.Select("write_keys").
		Columns("write_key", "owner_id", "source_id", "created_at", "last_used_at", "status").
		Where(qb.Eq("write_key")).
		ToCql()

	m := map[string]interface{}{
		"write_key": writeKey,
	}

	// Execute the query
	q := s.scyllaAdapter.ScyllaConn.Query(stmt, names)
	qx := q.BindMap(m)
	if err := qx.Get(&metadata); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return entity.WriteKeyMetaData{}, ErrWriteKeyNotFound
		}
		return entity.WriteKeyMetaData{}, err
	}

	return metadata, nil
}
