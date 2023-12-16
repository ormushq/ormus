package uservalidator

import "github.com/ormushq/ormus/manager/entity"

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type Repository interface {
	GetUserByEmail(email string) (*entity.User, error)
	DoesUserExistsByEmail(email string) (bool, error)
}
type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
