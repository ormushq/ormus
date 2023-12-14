package auth

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/richerror"
	"time"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	user := entity.User{
		ID:        "0",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		Email:     req.Email,
		Password:  req.Password,
		IsActive:  false,
	}
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, richerror.New("register").WhitWarpError(err)
	}
	//return create new user
	return param.RegisterResponse{
		Email: createdUser.Email,
	}, nil

}
