package userhandler

import (
	"github.com/ormushq/ormus/manager/validator/uservalidator"
	"github.com/ormushq/ormus/param"
)

type UserService interface {
	Login(req param.LoginRequest) (param.LoginResponse, error)
	Register(req param.RegisterRequest) (param.RegisterResponse, error)
}

type UserValidator interface {
	ValidateLoginRequest(req param.LoginRequest) *uservalidator.ValidatorError
	ValidateRegisterRequest(req param.RegisterRequest) *uservalidator.ValidatorError
}

type Handler struct {
	// TODO - add configurations
	userSvc       UserService
	userValidator UserValidator
}

func New(userSvc UserService, userValidator UserValidator) *Handler {
	return &Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
