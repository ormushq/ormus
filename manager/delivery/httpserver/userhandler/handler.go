package userhandler

import (
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/service/userservice"
)

type Handler struct {
	// TODO - add configurations
	userSvc    userservice.Service
	projectSvc projectservice.Service
}

func New(userSvc userservice.Service, projectSvc projectservice.Service) Handler {
	return Handler{
		userSvc:    userSvc,
		projectSvc: projectSvc,
	}
}
