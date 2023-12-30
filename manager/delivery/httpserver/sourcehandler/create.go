package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/param"
)

// ? Handler or *Handler.
func (h Handler) CreateSource(ctx echo.Context) error {
	// TODO  get owner(user) id

	// TODO  get project id ?

	// binding addsource request form
	AddSourceReq := new(param.AddSourceRequest)
	if err := ctx.Bind(AddSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	// validate form also check existen
	if err := h.validateSvc.ValidateCreateSourceForm(*AddSourceReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	// call save method in service
	sourceResp, err := h.sourceSvc.CreateSource(AddSourceReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	return ctx.JSON(http.StatusCreated, sourceResp) // TODO maybe need change response structure
}
