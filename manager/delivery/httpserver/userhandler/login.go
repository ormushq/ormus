package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/echomsg"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
)

// TODO: the naming convention is weird UserLogin, and RegisterUser should we change them?

func (h Handler) UserLogin(ctx echo.Context) error {
	var req param.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	result := h.userValidator.ValidateLoginRequest(req)
	if result != nil {
		msg, code := httpmsg.Error(result.Err)
		if result.Fields["email"] == "user not found" {
			return ctx.JSON(http.StatusUnauthorized, echomsg.DefaultMessage(errmsg.ErrWrongCredentials))
		}

		// TODO: in validator we have a ValidatorError struct and this binding is ambiguous, we should change it properly
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  result.Fields,
		})
	}

	response, err := h.userSvc.Login(req)
	if err != nil {
		if err.Error() == errmsg.ErrWrongCredentials {
			return ctx.JSON(http.StatusUnauthorized, echomsg.DefaultMessage(errmsg.ErrWrongCredentials))
		}

		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	// TODO : set cookie

	return ctx.JSON(http.StatusOK, response)
}
