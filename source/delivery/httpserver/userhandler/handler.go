package userhandler

import (
	"github.com/ormushq/ormus/manager/service/auth"
)

type Handler struct {
	// TODO - add configurations
	userSvc auth.Service
}

func New(userSvc auth.Service) *Handler {

	return &Handler{userSvc: userSvc}

}
