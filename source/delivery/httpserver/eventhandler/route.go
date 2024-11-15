package eventhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/delivery/middlewares"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	e.POST("api/source/event/", h.NewEvent, middlewares.WriteKeyMiddleware(h.eventValidator))

	// TODO - add required routes
}
