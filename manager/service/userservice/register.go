package userservice

import (
	"github.com/ormushq/ormus/param"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	// TODO : creat user in repository
	//return create new user
	return param.RegisterResponse{
		ID:    "0",
		Email: req.Email,
	}, nil

}
