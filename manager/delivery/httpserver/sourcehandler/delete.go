package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) DeleteSource(ctx echo.Context) error {
	// get id from param
	sourceID := ctx.Param("sourceId")

	// TODO validate id and check existen source ?

	// call delete service method
	if err := h.sourceSvc.DeleteSource(sourceID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	return ctx.JSON(http.StatusNoContent, echo.Map{"result": "delete"}) // TODO maybe need change response structure
}
