package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/pkg/echomsg"
)

// healthCheck godoc
//
//	@Summary		Show health check
//	@Description	get service health check
//	@Tags			healthCheck
//	@Accept			json
//	@Produce		json
//	@Router			/health-check [get]
func (s *Server) healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echomsg.DefaultMessage("health check message from server"))
}
