package scyllasource

import (
	"github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["GetWithId"] = scyllarepo.Statement{
		Query:  "SELECT id,token(id) as token_id, write_key, name, description, project_id, owner_id, status, created_at, updated_at, deleted_at FROM sources where id = ? and deleted = false  ALLOW FILTERING;",
		Values: []string{"id"},
	}
	statements["IsWriteKeyValid"] = scyllarepo.Statement{
		Query:  "SELECT write_key FROM sources where write_key = ? ALLOW FILTERING;",
		Values: []string{"write_key"},
	}
}

func (r Repository) GetWithID(id string) (entity.Source, error) {
	query, err := r.db.GetStatement(statements["GetWithId"])
	if err != nil {
		return entity.Source{}, err
	}

	query.BindMap(qb.M{
		"id": id,
	})

	var source entity.Source
	if err := query.Get(&source); err != nil {
		return entity.Source{}, err
	}

	return source, nil
}

func (r Repository) IsWriteKeyValid(writeKey string) (*source.ValidateWriteKeyResp, error) {
	query, err := r.db.GetStatement(statements["IsWriteKeyValid"])
	if err != nil {
		return nil, err
	}
	query.BindMap(
		qb.M{
			"write_key": writeKey,
		})
	var key string
	if err := query.Get(&key); err != nil {
		return nil, err
	}
	var writeKeyResp source.ValidateWriteKeyResp
	if key != "" {
		writeKeyResp.IsValid = true
	} else {
		writeKeyResp.IsValid = false
	}

	return &writeKeyResp, nil
}
