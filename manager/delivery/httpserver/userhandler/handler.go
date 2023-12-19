package userhandler

import (
	"github.com/ormushq/ormus/manager/service/userservice"
	uservalidator "github.com/ormushq/ormus/manager/validator/user"
)

type Handler struct {
	// TODO - add configurations
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(userSvc userservice.Service, userValidator uservalidator.Validator) *Handler {
	return &Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
