package statushandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/source/params"
)

func (h Handler) status(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, params.StatusResponse{
		Message: "System is functioning correctly.",
	})
}
