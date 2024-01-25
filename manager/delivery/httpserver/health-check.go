package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, EchoErrorMessage("health check message from server"))
}
