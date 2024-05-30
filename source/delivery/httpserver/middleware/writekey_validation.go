package middleware

import (
	"github.com/labstack/echo/v4"
)

func WritekeyValidation() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			//TODO - implement me
			return next(ctx)
		}
	}
}

func EchoErrorMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
