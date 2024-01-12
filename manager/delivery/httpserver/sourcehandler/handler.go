package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/delivery/httpserver/middleware"
	"github.com/ormushq/ormus/manager/service/sourceservice"
	"github.com/ormushq/ormus/manager/service/userservice"
	"github.com/ormushq/ormus/manager/validator/sourcevalidator"
)

type Handler struct {
	sourceSvc        sourceservice.Service
	userSvc          userservice.Service
	validateSvc      sourcevalidator.Validator
	customMiddleware middleware.Middleware
}

func New(sourceSvc sourceservice.Service,
	userSvc userservice.Service,
	validateSvc sourcevalidator.Validator,
	customMiddleware middleware.Middleware,
) *Handler {
	return &Handler{
		sourceSvc:        sourceSvc,
		userSvc:          userSvc,
		validateSvc:      validateSvc,
		customMiddleware: customMiddleware,
	}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
