package httpserver

import (
	"github.com/ormushq/ormus/pkg/echomsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echomsg.DefaultMessage("health check message from server"))
}
