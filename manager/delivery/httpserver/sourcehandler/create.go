package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/param"
)

// ? Handler or *Handler.
func (h Handler) CreateSource(ctx echo.Context) error {
	// get user id from context
	u := ctx.Get("userID")
	userID, ok := u.(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage("can not get userID"))
	}

	// TODO  get project id  if get from param dont forget add to route ?

	// binding addsource request form
	AddSourceReq := new(param.AddSourceRequest)
	if err := ctx.Bind(AddSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	// validate form also check existen
	if err := h.validateSvc.ValidateCreateSourceForm(*AddSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	// call save method in service
	sourceResp, err := h.sourceSvc.CreateSource(AddSourceReq, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, EchoErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, sourceResp)
}
