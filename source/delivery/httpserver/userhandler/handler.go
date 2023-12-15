package userhandler

import (
	service "github.com/ormushq/ormus/manager/service/auth"
)

type Handler struct {
	// TODO - add configurations
	userSvc service.Service
}

func New(userSvc service.Service) *Handler {
	return &Handler{userSvc: userSvc}
}
