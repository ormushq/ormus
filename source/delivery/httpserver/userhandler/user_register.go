package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/param"
	"net/http"
)

func (h *Handler) RegisterUser(ctx echo.Context) error {
	var Req param.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	resp, err := h.userSvc.Register(Req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	return ctx.JSON(http.StatusCreated, resp)
}
