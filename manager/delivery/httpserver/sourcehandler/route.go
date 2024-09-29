package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/delivery/httpserver/middleware"
)

func (h Handler) SetSourceRoute(e *echo.Echo) {
	sourceGroups := e.Group("/sources", middleware.GetTokenFromHeader(h.authSvc))
	sourceGroups.POST("/", h.Create)
	sourceGroups.POST("/:sourceID", h.Update)
	sourceGroups.DELETE("/:sourceID", h.Delete)
}
