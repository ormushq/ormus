package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/param"
)

func (h Handler) UpdateSource(ctx echo.Context) error {
	// TODO  get owner(user) id

	// TODO  get project id ?

	// get source id
	sourceID := ctx.Param("sourceID")

	// binding addsource request form
	updateSourceReq := new(param.UpdateSourceRequest)
	if err := ctx.Bind(updateSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	if err := h.validateSvc.ValidateUpdateSourceForm(*updateSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	// call save method in service
	sourceResp, err := h.sourceSvc.UpdateSource(sourceID, updateSourceReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	return ctx.JSON(http.StatusCreated, sourceResp) // TODO maybe need change response structure
}
