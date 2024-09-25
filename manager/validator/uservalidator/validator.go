package uservalidator

import (
	"github.com/ormushq/ormus/manager/entity"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	DoesUserExistsByEmail(email string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
