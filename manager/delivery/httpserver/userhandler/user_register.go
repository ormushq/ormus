package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/httpmsg"
)

func (h Handler) RegisterUser(ctx echo.Context) error {
	var Req param.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	result := h.userValidator.ValidateRegisterRequest(Req)
	if result != nil {
		msg, code := httpmsg.Error(result.Error)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"error":   result.Fields,
		})
	}

	resp, err := h.userSvc.Register(Req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	return ctx.JSON(http.StatusCreated, resp)
}
