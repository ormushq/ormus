package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/sourceservice"
	"github.com/ormushq/ormus/manager/validator/sourcevalidator"
)

type Handler struct {
	sourceSvc   sourceservice.Service
	ValidateSvc sourcevalidator.Validator
	// TODO all serveices
}

func New() *Handler { // TODO give parameters
	return &Handler{}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
