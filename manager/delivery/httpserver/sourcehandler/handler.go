package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/service/sourceservice"
)

type Handler struct {
	sourceSvc sourceservice.Service
	authSvc   authservice.Service
}

func New(authSvc authservice.Service, sourceSvc sourceservice.Service) Handler {
	return Handler{
		sourceSvc: sourceSvc,
		authSvc:   authSvc,
	}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
