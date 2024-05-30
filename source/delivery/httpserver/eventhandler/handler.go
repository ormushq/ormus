package eventhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/service/eventservice"
	"github.com/ormushq/ormus/source/validator/eventvalidator"
)

type Handler struct {
	eventSvc eventservice.Service
	eventVld eventvalidator.Validator
}

func New(eventSvc eventservice.Service,
	eventVld eventvalidator.Validator,
) Handler {
	return Handler{
		eventSvc: eventSvc,
		eventVld: eventVld,
	}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
