package auth

import "github.com/ormushq/ormus/manager/entity"

type Repository interface {
	Register(u entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type JWTEngine interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	repo Repository
	jwt  JWTEngine
}

func NewService(authGenerator JWTEngine, repository Repository) Service {
	return Service{jwt: authGenerator, repo: repository}
}
