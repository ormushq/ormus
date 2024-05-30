package userhandler

import (
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
)

type Handler struct {
	// TODO - add configurations
	userSvc       *userservice.Service
	userValidator uservalidator.Validator
	projectSvc    *projectservice.Service
}

func New(userSvc *userservice.Service, userValidator uservalidator.Validator, projectSvc *projectservice.Service) *Handler {
	return &Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
		projectSvc:    projectSvc,
	}
}
