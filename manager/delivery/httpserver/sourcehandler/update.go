package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/managerparam"
)

func (h Handler) UpdateSource(ctx echo.Context) error {
	// get user id from context
	u := ctx.Get("userID")
	userID, ok := u.(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage("can not get userID"))
	}

	// TODO  get project id ?

	// get source id
	sourceID := ctx.Param("sourceID")

	// binding addsource request form
	updateSourceReq := new(managerparam.UpdateSourceRequest)
	if err := ctx.Bind(updateSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	if err := h.validateSvc.ValidateUpdateSourceForm(*updateSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	// call save method in service
	sourceResp, err := h.sourceSvc.UpdateSource(userID, sourceID, updateSourceReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, sourceResp)
}
