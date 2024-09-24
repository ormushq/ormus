package statushandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/adapter/otela"
	"github.com/ormushq/ormus/source/params"
)

func (h Handler) status(ctx echo.Context) error {
	// You can pass newCtx to service for set this span as parent of spans that start in services
	_, span := otela.TraceBuilder("statushandler", "status",
		otela.WithContext(ctx.Request().Context()),
	)
	defer span.End()

	span.AddEvent("Starting status handler")
	// Do something
	span.AddEvent("Ending status handler")

	return ctx.JSON(http.StatusOK, params.StatusResponse{
		Message: "System is functioning correctly.",
	})
}
