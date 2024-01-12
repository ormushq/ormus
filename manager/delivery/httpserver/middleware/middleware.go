package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/service/authservice"
)

type Middleware struct {
	jwtSvc authservice.JWT
}

func New(jwtSvc authservice.JWT) *Middleware {
	return &Middleware{
		jwtSvc: jwtSvc,
	}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
