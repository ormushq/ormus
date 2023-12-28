package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/sourceservice"
	"github.com/ormushq/ormus/manager/validator/sourcevalidator"
)

type Handler struct {
	sourceSvc   sourceservice.Service
	validateSvc sourcevalidator.Validator
}

func New(sourceSvc sourceservice.Service,
	validateSvc sourcevalidator.Validator,
) *Handler {
	return &Handler{sourceSvc: sourceSvc, validateSvc: validateSvc}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
