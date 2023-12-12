package userservice

import "github.com/ormushq/ormus/manager/entity"

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	// TODO: implement repository
	//repo Repository
	auth AuthGenerator
}

func New(authGenerator AuthGenerator) Service {
	return Service{auth: authGenerator}
}
