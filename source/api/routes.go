package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all routes for the API.
func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Development mode.")
    })
	
	// Add more routes as needed
}