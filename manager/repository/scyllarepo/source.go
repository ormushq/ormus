package scyllarepo

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/scylladb/gocqlx/v2/qb"
)

func (a StorageAdapter) InsertSource(source *entity.Source) (*managerparam.AddSourceResponse, error) {
	stmt, names := qb.Insert("sources").
		Columns("id", "writekey", "name", "description", "owner_id", "project_id").ToCql()
	q := a.ScyllaConn.Query(stmt, names)
	err := q.Bind(source).ExecRelease()
	if err != nil {
		logger.L().Error("Failed to insert source", "error", err)

		return nil, err
	}

	return &managerparam.AddSourceResponse{
		ID:          source.ID,
		WriteKey:    string(source.WriteKey),
		Name:        source.Name,
		Description: source.Description,
		OwnerID:     source.OwnerID,
		ProjectID:   source.ProjectID,
	}, nil
}

func (a StorageAdapter) UpdateSource(id string, source *entity.Source) (*managerparam.UpdateSourceResponse, error) {
	stmt, names := qb.Update("sources").
		Set("name", "description").
		Where(qb.Eq("id")).
		ToCql()
	q := a.ScyllaConn.Query(stmt, names)
	err := q.BindMap(qb.M{
		"id":          id,
		"name":        source.Name,
		"description": source.Description,
	}).ExecRelease()
	if err != nil {
		logger.L().Error("Failed to update source", "error", err)

		return nil, err
	}

	return &managerparam.UpdateSourceResponse{
		ID:          id,
		WriteKey:    string(source.WriteKey),
		Name:        source.Name,
		Description: source.Description,
		ProjectID:   source.ProjectID,
		OwnerID:     source.OwnerID,
	}, nil
}

func (a StorageAdapter) DeleteSource(id, userID string) error {
	stmt, names := qb.Delete("sources").
		Where(qb.Eq("id"), qb.Eq("owner_id")).
		ToCql()
	q := a.ScyllaConn.Query(stmt, names)
	err := q.BindMap(qb.M{
		"id":       id,
		"owner_id": userID,
	}).ExecRelease()
	if err != nil {
		logger.L().Error("Failed to delete source", "error", err)

		return err
	}

	return nil
}

func (a StorageAdapter) GetUserSourceByID(ownerID, id string) (*entity.Source, error) {
	stmt, names := qb.Select("sources").
		Columns("id", "writekey", "name", "description", "owner_id", "project_id").
		Where(qb.Eq("id"), qb.Eq("owner_id")).
		ToCql()
	q := a.ScyllaConn.Query(stmt, names)
	qx := q.BindMap(qb.M{
		"id":       id,
		"owner_id": ownerID,
	})
	var source entity.Source
	if err := qx.GetRelease(&source); err != nil {
		logger.L().Error("Failed to get source", "error", err)

		return nil, err
	}

	return &source, nil
}

func (a StorageAdapter) IsSourceAlreadyCreatedByName(name string) (bool, error) {
	stmt, names := qb.Select("sources").
		Columns("COUNT(1)").
		Where(qb.Eq("name")).
		ToCql()
	q := a.ScyllaConn.Query(stmt, names)
	qx := q.BindMap(qb.M{
		"name": name,
	})
	var count int
	if err := qx.GetRelease(&count); err != nil {
		logger.L().Error("Failed to check if source already created by name", "error", err)

		return false, err
	}

	return count > 0, nil
}
