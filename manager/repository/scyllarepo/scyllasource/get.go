package scyllasource

import (
	source_proto "github.com/ormushq/ormus/contract/go/source"
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
		Query:  "SELECT write_key,project_id,owner_id FROM sources where write_key = ? ALLOW FILTERING;",
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

func (r Repository) IsWriteKeyValid(writeKey string) (*source_proto.ValidateWriteKeyResp, error) {
	query, err := r.db.GetStatement(statements["IsWriteKeyValid"])
	if err != nil {
		return nil, err
	}
	query.BindMap(
		qb.M{
			"write_key": writeKey,
		})
	var writeKeyResp source_proto.ValidateWriteKeyResp
	if err := query.Get(&writeKeyResp); err != nil && err.Error() != "not found" {
		return nil, err
	}

	if writeKeyResp.WriteKey != "" {
		writeKeyResp.IsValid = true
	} else {
		writeKeyResp.IsValid = false
	}

	return &writeKeyResp, nil
}
