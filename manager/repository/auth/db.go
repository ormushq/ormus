package repository

import (
	"fmt"

	"github.com/ormushq/ormus/manager/entity"
)

// TODO: implement repository for auth

type StorageAdapter struct{}

func (a StorageAdapter) Register(u entity.User) (*entity.User, error) {
	fmt.Println(u)
	panic("implement me")
}

func (a StorageAdapter) GetUserByEmail(email string) (*entity.User, error) {
	fmt.Println(email)
	panic("implement me")
}
