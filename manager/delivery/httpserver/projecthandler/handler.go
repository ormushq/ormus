package projecthandler

import (
	"github.com/ormushq/ormus/manager/service/projectservice"
	"github.com/ormushq/ormus/manager/validator/projectvalidator"
)

type Handler struct {
	projectSvc       *projectservice.Service
	projectValidator projectvalidator.Validator
}

func New(projectSvc *projectservice.Service, projectValidator projectvalidator.Validator) *Handler {
	return &Handler{
		projectSvc:       projectSvc,
		projectValidator: projectValidator,
	}
}
