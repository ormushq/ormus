package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/delivery/httpserver/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	sourceGroups := e.Group("/sources", middleware.GetTokenFromHeader(h.authSvc))
	sourceGroups.GET("", h.List)
	sourceGroups.POST("", h.Create)
	sourceGroups.POST("/:sourceID", h.Update)
	sourceGroups.DELETE("/:sourceID", h.Delete)
}
