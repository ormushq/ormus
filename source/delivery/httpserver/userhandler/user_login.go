package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/param"
)

func (h Handler) UserLogin(ctx echo.Context) error {
	var req param.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	response, err := h.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, response)
}
