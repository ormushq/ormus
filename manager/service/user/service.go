package user

import "github.com/ormushq/ormus/manager/entity"

type Repository interface {
	Register(u entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	repo Repository
	auth AuthGenerator
}

func New(authGenerator AuthGenerator) Service {
	return Service{auth: authGenerator}
}
