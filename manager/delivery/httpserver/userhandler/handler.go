package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/uservalidator"
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

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
