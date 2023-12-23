package userservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/password"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Register(req param.RegisterRequest) (*param.RegisterResponse, error) {
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, richerror.New("register.hash").WhitWarpError(err)
	}

	user := entity.User{
		DeletedAt: nil,
		Email:     req.Email,
		Password:  hashedPassword,
		IsActive:  false,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, richerror.New("register.repo").WhitWarpError(err)
	}

	// return create new user
	return &param.RegisterResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}, nil
}
