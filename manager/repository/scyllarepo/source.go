package scyllarepo

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
)

func (a StorageAdapter) InsertSource(source *entity.Source) (*managerparam.AddSourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a StorageAdapter) UpdateSource(id string, source *entity.Source) (*managerparam.UpdateSourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a StorageAdapter) DeleteSource(id, userID string) error {
	//TODO implement me
	panic("implement me")
}

func (a StorageAdapter) GetUserSourceByID(ownerID, id string) (*entity.Source, error) {
	//TODO implement me
	panic("implement me")
}

func (a StorageAdapter) IsSourceAlreadyCreatedByName(name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
