package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoute(e *echo.Echo) {

	userGroups := e.Group("/users")
	userGroups.POST("/register", h.RegisterUser)
}
