package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "health check message from server",
	})
}
