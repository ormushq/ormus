package eventhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/pkg/echomsg"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
	"github.com/ormushq/ormus/source/params"
	"net/http"
)

// ? Handler or *Handler.
func (h Handler) TrackEvent(ctx echo.Context) error {
	var req params.TrackEventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	result := h.eventVld.ValidateTrackRequest(req)
	if result != nil {
		msg, code := httpmsg.Error(result.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  result.Fields,
		})
	}

	resp, err := h.eventSvc.Track(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	return ctx.JSON(http.StatusCreated, resp)

}
