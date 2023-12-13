package userhandler

import (
	"github.com/ormushq/ormus/manager/service/user"
)

type Handler struct {
	// TODO - add configurations
	userSvc user.Service
}

func New(userSvc user.Service) *Handler {

	return &Handler{userSvc: userSvc}

}
