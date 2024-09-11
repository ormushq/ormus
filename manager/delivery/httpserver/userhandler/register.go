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

// RegisterUser godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		param.RegisterRequest	true	"Register request body"
//	@Success		201		{object}	param.RegisterResponse
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/users/register [post]
func (h Handler) RegisterUser(ctx echo.Context) error {
	var Req param.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	// TODO: should we return service error? or should we only return bad request error?
	resp, err := h.userSvc.Register(Req)

	var vErr *validator.Error
	if errors.As(err, &vErr) {
		msg, code := httpmsg.Error(vErr.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	return ctx.JSON(http.StatusCreated, resp)
}
