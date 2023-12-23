package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Expose routes for server connections.
func (h Handler) SetRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Development mode.")
	})

	// TODO - add required routes
}
