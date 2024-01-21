package sourcehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/delivery/httpserver/middleware"
)

func (h Handler) SetSourceRoute(e *echo.Echo) {
	sourceGroups := e.Group("/source")
	sourceGroups.POST("/create", h.CreateSource, middleware.GetTokenFromCookie(h.authSvc))
	sourceGroups.POST("/update/:sourceID", h.UpdateSource, middleware.GetTokenFromCookie(h.authSvc))
	sourceGroups.DELETE("/remove/:sourceID", h.DeleteSource)
}
