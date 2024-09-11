package userhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
	"github.com/ormushq/ormus/pkg/httputil"
)

// TODO: the naming convention is weird UserLogin, and RegisterUser should we change them?

// UserLogin godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		param.LoginRequest	true	"Login request body"
//	@Success		200		{object}	param.LoginResponse
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		401		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/users/login [post]
func (h Handler) UserLogin(ctx echo.Context) error {
	var req param.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	response, err := h.userSvc.Login(req)

	var vErr *validator.Error
	if errors.As(err, &vErr) {
		msg, code := httpmsg.Error(vErr.Err)
		if vErr.Fields["email"] == "user not found" {
			return httputil.NewError(ctx, http.StatusUnauthorized, errmsg.ErrWrongCredentials)
		}

		// TODO: in validator we have a ValidatorError struct and this binding is ambiguous, we should change it properly
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}
	if err != nil {
		if err.Error() == errmsg.ErrWrongCredentials {
			return httputil.NewError(ctx, http.StatusUnauthorized, errmsg.ErrWrongCredentials)
		}

		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	// TODO : set cookie

	return ctx.JSON(http.StatusOK, response)
}
