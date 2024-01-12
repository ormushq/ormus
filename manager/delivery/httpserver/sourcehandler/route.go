package sourcehandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetSourceRoute(e *echo.Echo) {
	sourceGroups := e.Group("/source")
	sourceGroups.POST("/create", h.CreateSource, h.customMiddleware.GetTokenFromCookie)
	sourceGroups.POST("/update/:sourceID", h.UpdateSource, h.customMiddleware.GetTokenFromCookie)
	sourceGroups.DELETE("/remove/:sourceID", h.DeleteSource)
}
