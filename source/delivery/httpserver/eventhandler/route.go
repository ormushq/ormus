package eventhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/delivery/httpserver/middleware"
)

func (h Handler) SetEventRoute(e *echo.Echo) {
	eventGroups := e.Group("/v1/track-event")
	eventGroups.POST("/", h.TrackEvent, middleware.WritekeyValidation())
}
