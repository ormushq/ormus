package statushandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(router *echo.Echo) {
	router.GET("api/source/status", h.status)
}