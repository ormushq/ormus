package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Delete(ctx echo.Context) error {
	// get user id from context
	u := ctx.Get("userID")
	userID, ok := u.(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage("can not get userID"))
	}
	// get id from param
	sourceID := ctx.Param("sourceID")

	// validate form also check existen
	if err := h.validateSvc.ValidateIDToDelete(sourceID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	// call delete service method
	if err := h.sourceSvc.DeleteSource(sourceID, userID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
