package projecthandler

import (
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/projectservice"
)

type Handler struct {
	projectSvc projectservice.Service
	authSvc    authservice.Service
}

func New(authSvc authservice.Service, projectSvc projectservice.Service) Handler {
	return Handler{
		projectSvc: projectSvc,
		authSvc:    authSvc,
	}
}
