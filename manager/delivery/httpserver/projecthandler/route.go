package projecthandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoute(e *echo.Echo) {
	projectGroup := e.Group("/project")

	projectGroup.POST("/", h.Create)
}
