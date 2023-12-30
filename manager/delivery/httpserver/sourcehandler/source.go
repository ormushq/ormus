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

func (h Handler) UpdateSource(ctx echo.Context) error {
	// TODO  get owner(user) id

	// TODO  get project id ?

	// get source id
	sourceID := ctx.Param("sourceId")

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

func (h Handler) DeleteSource(ctx echo.Context) error {
	// get id from param
	sourceID := ctx.Param("sourceId")

	// TODO validate id and check existen source ?

	// call delete service method
	if err := h.sourceSvc.DeleteSource(sourceID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error())) // TODO maybe need change response structure
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{"result": "delete"}) // TODO maybe need change response structure
}
