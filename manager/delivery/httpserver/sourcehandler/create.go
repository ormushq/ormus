package sourcehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/param"
)

// ? Handler or *Handler.
func (h Handler) CreateSource(ctx echo.Context) error {
	// TODO  get owner(user) id
	// to implement this todo
	//		1 : give token from header or cookie but in login handler just return jwt token as response body
	// 		2 : then pars token with jwt.ParsToken() function i think its better that we implement middelware to pars and set email in contxt "ctx.set(...)"
	// 		3 : than here "ctx.Get(...)" than pass email to service layer to find owner(user) id to create new source

	// TODO  get project id ?

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
	sourceResp, err := h.sourceSvc.CreateSource(AddSourceReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, EchoErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, sourceResp)
}
