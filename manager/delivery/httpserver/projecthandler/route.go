package projecthandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/delivery/httpserver/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	projectGroups := e.Group("/projects", middleware.GetTokenFromHeader(h.authSvc))
	projectGroups.GET("", h.List)
	projectGroups.POST("", h.Create)
	projectGroups.POST("/:projectID", h.Update)
	projectGroups.DELETE("/:projectID", h.Delete)
}
