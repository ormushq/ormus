package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/httpmsg"
)

// TODO: the naming convention is weird UserLogin, and RegisterUser should we change them?

func (h Handler) UserLogin(ctx echo.Context) error {
	var req param.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	result := h.userValidator.ValidateLoginRequest(req)
	if result != nil {
		msg, code := httpmsg.Error(result.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"error":   result.Fields,
		})
	}

	response, err := h.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, response)
}
