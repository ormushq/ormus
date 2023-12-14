package repository

import "github.com/ormushq/ormus/manager/entity"

type StorageAdapter struct {
}

func (a StorageAdapter) Register(u entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a StorageAdapter) GetUserByEmail(email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
