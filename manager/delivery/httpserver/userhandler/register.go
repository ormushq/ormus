package userhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/echomsg"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
)

func (h Handler) RegisterUser(ctx echo.Context) error {
	var Req param.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	// TODO: should we return service error? or should we only return bad request error?
	resp, err := h.userSvc.Register(Req)

	vErr := validator.Error{}
	if errors.Is(err, vErr) {
		msg, code := httpmsg.Error(vErr.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	return ctx.JSON(http.StatusCreated, resp)
}
